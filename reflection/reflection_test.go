package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}
type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two field",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string fields",
			struct {
				Name string
				Age  int
			}{"Chris", 28},
			[]string{"Chris"},
		},
		{
			"struct with nested fields",
			Person{"Chris", Profile{
				28, "London",
			}},
			[]string{"Chris", "London"},
		},
		{
			"pointer to things",
			&Person{"Chris", Profile{
				28, "London",
			}},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{{
				28, "London",
			}, {
				29, "Tokyo",
			}},
			[]string{"London", "Tokyo"},
		},
		{
			"arrays",
			[2]Profile{{
				28, "London",
			}, {
				29, "Tokyo",
			}},
			[]string{"London", "Tokyo"},
		},
		{
			"maps",
			map[string]string{
				"Cow":   "Moo",
				"Sheep": "bee",
			},
			[]string{"Moo", "bee"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
	t.Run("with Maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})
	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{88, "Berlin"}
			aChannel <- Profile{12, "Warszawa"}
			close(aChannel)
		}()
		var got []string
		want := []string{"Berlin", "Warszawa"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{88, "Berlin"}, Profile{12, "Warszawa"}
		}
		var got []string
		want := []string{"Berlin", "Warszawa"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}

package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	// Output: aaaaa
}

func TestRepeat(t *testing.T) {
	t.Run("repeat with custom function", func(t *testing.T) {
		repeated := Repeat("a", 8)
		expected := "aaaaaaaa"
		assertCorrectOutput(t, repeated, expected)
	})
	t.Run("repeat comopare with standard library built-in", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := strings.Repeat("a", 5)
		assertCorrectOutput(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 8)
	}
}

func assertCorrectOutput(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

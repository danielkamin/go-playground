package main

import "testing"

func TestSearch(t *testing.T) {
	word := "this is just a test"
	key := "test"
	dictionary := Dictionary{key: word}
	t.Run("known word", func(t *testing.T) {
		assertDefinition(t, dictionary, key, word)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("xd")
		if got == nil {
			t.Fatal("expected to get an error.")
		}
		assertErrors(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Adding new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "this is just a test"
		key := "test"
		dictionary.Add(key, word)
		assertDefinition(t, dictionary, key, word)
	})
	t.Run("Adding existing word", func(t *testing.T) {
		word := "this is just a test"
		key := "test"
		dictionary := Dictionary{key: word}
		err := dictionary.Add(key, "new word")
		assertErrors(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, word)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update existing word", func(t *testing.T) {
		word := "this is just a test"
		key := "test"
		dictionary := Dictionary{key: word}
		newWord := "this is a test updated"
		err := dictionary.Update(key, newWord)
		assertErrors(t, err, nil)
		assertDefinition(t, dictionary, key, newWord)
	})
	t.Run("Update new word", func(t *testing.T) {
		word := "this is just a test"
		key := "test"
		dictionary := Dictionary{}
		err := dictionary.Update(key, word)
		assertErrors(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {

	t.Run("Delete existing word", func(t *testing.T) {
		word := "this is just a test"
		key := "test"
		dictionary := Dictionary{key: word}
		err := dictionary.Delete(key)
		assertErrors(t, err, nil)
		_, err = dictionary.Search(key)
		assertErrors(t, err, ErrNotFound)
	})
	t.Run("Delete not existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete("another key")
		assertErrors(t, err, ErrWordDoesNotExist)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
func assertErrors(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
func assertDefinition(t testing.TB, dictionary Dictionary, key, word string) {
	t.Helper()
	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("Should find added word: ", err)
	}
	assertStrings(t, got, word)
}

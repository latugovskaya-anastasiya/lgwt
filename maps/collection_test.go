package collection_test

import (
	"testing"

	collection "github.com/latugovskaya-anastasiya/lgwt/maps"
)

func TestSearch(t *testing.T) {
	dictionary := collection.Dictionary{"test": "this is just a test"}

	t.Run(
		"known word", func(t *testing.T) {
			got, _ := dictionary.Search("test")
			want := "this is just a test"

			assertStrings(t, got, want)
		},
	)

	t.Run(
		"unknown word", func(t *testing.T) {
			_, got := dictionary.Search("unknown")

			assertError(t, got, collection.ErrNotFound)
		},
	)
}

func TestAdd(t *testing.T) {
	t.Run(
		"new word", func(t *testing.T) {
			dictionary := collection.Dictionary{}
			word := "test"
			definition := "this is just a test"

			err := dictionary.Add(word, definition)

			assertError(t, err, nil)
			assertDefinition(t, dictionary, word, definition)
		},
	)

	t.Run(
		"existing word", func(t *testing.T) {
			word := "test"
			definition := "this is just a test"
			dictionary := collection.Dictionary{word: definition}
			err := dictionary.Add(word, "new test")

			assertError(t, err, collection.ErrWordExists)
			assertDefinition(t, dictionary, word, definition)
		},
	)
}

func TestUpdate(t *testing.T) {
	t.Run(
		"existing word", func(t *testing.T) {
			word := "test"
			definition := "this is just a test"
			newDefinition := "new definition"
			dictionary := collection.Dictionary{word: definition}
			err := dictionary.Update(word, newDefinition)

			assertError(t, err, nil)
			assertDefinition(t, dictionary, word, newDefinition)
		},
	)

	t.Run(
		"new word", func(t *testing.T) {
			word := "test"
			definition := "this is just a test"
			dictionary := collection.Dictionary{}

			err := dictionary.Update(word, definition)

			assertError(t, err, collection.ErrWordDoesNotExist)
		},
	)
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := collection.Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != collection.ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary collection.Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

package maps

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "test value"}

	t.Run("returns definitions for known words", func(t *testing.T) {
		assertDefinition(t, dictionary, "test", "test value")
	})

	t.Run("returns an error when an unknown word is searched for", func(t *testing.T) {
		result, actualError := dictionary.Search("unknown word")

		if result != "" {
			t.Fatal("should not return a definition")
		}

		assertError(t, actualError, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("can add new words", func(t *testing.T) {
		dictionary := Dictionary{}
		newWord := "new word"
		newDefinition := "new definition"

		dictionary.Add(newWord, newDefinition)

		assertDefinition(t, dictionary, newWord, newDefinition)
	})

	t.Run("errors if a word that already exists is added", func(t *testing.T) {
		existingWord := "existingWord"
		existingDefinition := "a definition for this word already exists"
		dictionary := Dictionary{existingWord: existingDefinition}
		newDefinition := "this is a new definition for an existing word"

		alreadyExistingWordError := dictionary.Add(existingWord, newDefinition)

		assertError(t, alreadyExistingWordError, ErrWordConflict)
		assertDefinition(t, dictionary, existingWord, existingDefinition)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("should update the definition for an existing word", func(t *testing.T) {
		existingWord := "existing word"
		existingDefinition := "initial definiton for word"
		dictionary := Dictionary{existingWord: existingDefinition}

		newDefinition := "new definition for the word"
		dictionary.Update(existingWord, newDefinition)

		assertDefinition(t, dictionary, existingWord, newDefinition)
	})

	t.Run("should error when updating a word that does not exist", func(t *testing.T) {
		dictionary := Dictionary{}

		newWord := "new word"
		updateNewWordError := dictionary.Update(newWord, "definition for new word")

		assertError(t, updateNewWordError, ErrNotFound)

		_, wordNotExistErr := dictionary.Search(newWord)
		assertError(t, wordNotExistErr, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("can delete a word that exists", func(t *testing.T) {
		existingWord := "existing word"
		dictionary := Dictionary{existingWord: "initial definiton for word"}

		dictionary.Delete(existingWord)

		_, searchErr := dictionary.Search(existingWord)
		assertError(t, searchErr, ErrNotFound)
	})

	t.Run("does not error when deleting a word that does not exist", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "word"
		dictionary.Delete(word)

		_, searchErr := dictionary.Search(word)
		assertError(t, searchErr, ErrNotFound)
	})
}

func assertDefinition(tb testing.TB, dictionary Dictionary, word, expectedDefinition string) {
	tb.Helper()

	actualDefinition, searchError := dictionary.Search(word)

	if searchError != nil {
		tb.Fatal(fmt.Printf("received unexpected error: %v", searchError))
	}

	if actualDefinition != expectedDefinition {
		tb.Errorf("got %v, expected %v", actualDefinition, expectedDefinition)
	}
}

func assertError(tb testing.TB, actual, expected error) {
	tb.Helper()

	if actual == nil {
		tb.Fatal("should receive an error")
	}

	if actual != expected {
		tb.Errorf("got error %v, expected %v", actual, expected)
	}
}

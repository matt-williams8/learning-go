package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to someone when their name is provided", func(t *testing.T) {
		got := Hello("Matt", "en")
		want := "Hello, Matt"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when empty name is provided", func(t *testing.T) {
		got := Hello("", "en")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in English if an unrecongised language is requested", func(t *testing.T) {
		got := Hello("Matt", "some-lang")
		want := "Hello, Matt"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in Spanish", func(t *testing.T) {
		got := Hello("Matt", "es")
		want := "Hola, Matt"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("Matt", "fr")
		want := "Bonjour, Matt"

		assertCorrectMessage(t, got, want)
	})
}

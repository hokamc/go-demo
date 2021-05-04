package main

import "testing"

func TestHello(t *testing.T) {
	assertMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to maggie", func(t *testing.T) {
		got := Hello("maggie")
		want := "Hello, maggie!"
		assertMessage(t, got, want)
	})

	t.Run("say hello to error", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world!"
		assertMessage(t, got, want)
	})
}

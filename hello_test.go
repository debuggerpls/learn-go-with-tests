package main

import "testing"

func TestHello(t *testing.T) {
	// t.Run() is used to group tests. This is helpful when using shared code
	// that is used by other tests to spare set up.
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Deividas", "English")
		want := "Hello, Deividas"

		assert(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assert(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assert(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Alban", "French")
		want := "Bonjour, Alban"

		assert(t, got, want)
	})
}

// testing.TB is an interface for both Tests and Benchmarks, so helper can be reused.
func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

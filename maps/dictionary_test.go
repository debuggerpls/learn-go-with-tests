package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"

		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("unknown workd", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrKeyNotFound

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		want := "this is just a test"
		dictionary.Add("test", want)

		got, err := dictionary.Search("test")
		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "existing test"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, "some new test")
		assertError(t, err, ErrKeyExists)
		got, err := dictionary.Search(key)
		assertNoError(t, err)
		assertStrings(t, got, value)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("expected no error, got %q", got)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

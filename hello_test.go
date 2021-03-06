package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello when empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"

		assertCorrectMessage(t, want, got)
	})

}

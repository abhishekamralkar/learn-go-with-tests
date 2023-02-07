package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Saying hello to people in hind", func(t *testing.T) {
		got := Hello("Anay", "hindi")
		want := "Namaste Anay"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hola to people in spanish", func(t *testing.T) {
		got := Hello("Anay", "spanish")
		want := "Hola Anay"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to people in english", func(t *testing.T) {
		got := Hello("Anay", "english")
		want := "Hello Anay"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

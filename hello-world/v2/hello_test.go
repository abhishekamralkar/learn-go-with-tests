package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Anay")
	want := "Hello Anay"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

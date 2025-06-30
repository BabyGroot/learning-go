package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Marty")
	want := "Hello, Marty"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

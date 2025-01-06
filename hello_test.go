package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if want != got {
		t.Errorf("got %q != want %q", got, want)
	}
}

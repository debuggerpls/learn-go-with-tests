package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Deividas")
	want := "Hello, Deividas"

	if want != got {
		t.Errorf("got %q != want %q", got, want)
	}
}

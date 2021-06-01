package main

import "testing"

func TestHello(t *testing.T) {

	got := Hello("Chai")
	want := "Hello, Chai"

	if got!=want{
		t.Errorf("got %q, want %q", got, want)
	}
}

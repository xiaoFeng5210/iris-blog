package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("xiaofeng")
	want := "hello, xiaofeng"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

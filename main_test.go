package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hallo(); got != want {
		t.Errorf("main() return unexpected")
	}
}

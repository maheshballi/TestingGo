package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to the world", func(t *testing.T) {
		got := Hello("world")
		want := "hello world"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("saying hello to the world", func(t *testing.T) {

		got := Hello("")
		want := "hello world"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}

	})
}

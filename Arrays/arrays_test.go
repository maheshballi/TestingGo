package main

import "testing"

func TestAdd(t *testing.T) {
	sum := [5]int{1, 2, 3, 4, 5}

	got := Add(sum)
	want := 15

	if got != want {
		t.Errorf("%q %q", got, want)
	}

}

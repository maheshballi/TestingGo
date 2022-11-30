package main

import "testing"

func TestAdd(t *testing.T) {

	sum := []int{2, 36, 7, 54, 76, 45}

	got := Add(sum)
	want := 220

	if got != want {
		t.Errorf("%q %q", got, want)
	}
}

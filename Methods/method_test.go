package main

import "testing"

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}

		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("%v%v", got, want)
		}
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}

	})

}

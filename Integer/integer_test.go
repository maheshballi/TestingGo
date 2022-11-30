package integer

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 2)
	want := 0

	if got != want {
		t.Errorf("%q%q", got, want)
	}
}

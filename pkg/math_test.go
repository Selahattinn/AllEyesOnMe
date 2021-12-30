package pkg

import (
	"testing"
)

func TestAdd(t *testing.T) {

	got := Add(4, 3)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

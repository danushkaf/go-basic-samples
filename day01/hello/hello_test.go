package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "99 bottles of beer on the wall, 99 bottles of beer, ..."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

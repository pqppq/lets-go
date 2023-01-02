package assert

import "testing"

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper() // Hello, I'm testing helper func!

	if actual != expected {
		t.Errorf("expected %v, but got %v instead.", expected, actual)
	}
}

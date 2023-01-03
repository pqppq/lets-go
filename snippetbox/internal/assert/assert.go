package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper() // Hello, I'm testing helper func!

	if actual != expected {
		t.Errorf("expected %v, but got %v instead.", expected, actual)
	}
}

func StringContains(t *testing.T, actual, expectedSubstring string) {
	t.Helper()

	if !strings.Contains(actual, expectedSubstring) {
		t.Errorf("extected %q, got %q instead.", expectedSubstring, actual)
	}
}

func NilError(t *testing.T, actual error) {
	t.Errorf("expected nil, got %v instead.", actual)
}

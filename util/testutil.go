package util

import (
	"fmt"
	"testing"
)

func Assert(t *testing.T, want any, got string) {
	if fmt.Sprint(want) != got {
		t.Fatalf("wanted '%s', got '%s'", want, got)
	}
}

func AssertEq[K comparable](t *testing.T, want K, got K) {
	if want != got {
		t.Fatalf("wanted '%s', got '%s'", want, got)
	}
}

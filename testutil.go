package main

import (
	"fmt"
	"testing"
)

func Assert(t *testing.T, want any, got string) {
	if fmt.Sprint(want) != got {
		t.Fatalf("wanted '%s', got '%s'", want, got)
	}
}

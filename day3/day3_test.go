package day3_test

import (
	"testing"
)

func TestPlaceholder(t *testing.T) {
	want := 1
	got := 1
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

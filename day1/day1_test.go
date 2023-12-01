package day1_test

import (
	"day1"
	"testing"
)

func TestFindsTheFirstAndLastNumber(t *testing.T) {
	t.Parallel()
	str := "1abc2"
	want := 12
	got := day1.parseLine(str)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

package day6

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLoadInputIntoTimeDistancePairs(t *testing.T) {
	input := "resources/input_test.txt"
	want := []Race{
		{
			Time:     7,
			Distance: 9,
		},
		{
			Time:     15,
			Distance: 40,
		},
		{
			Time:     30,
			Distance: 200,
		},
	}
	got := loadInput(input)
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestCheckWinScenarios(t *testing.T) {
	input := Race{
		Time:     7,
		Distance: 9,
	}
	want := 4
	got := checkWinScenarios(input)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSolvePart1(t *testing.T) {
	filePath := "resources/input_test.txt"
	want := 288
	got := Solve(filePath)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

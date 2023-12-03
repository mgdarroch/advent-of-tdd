package day3

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLoadInput(t *testing.T) {
	want := [][]string{
		{"4", "6", "7", ".", ".", "1", "1", "4", ".", "."},
		{".", ".", ".", "*", ".", ".", ".", ".", ".", "."},
	}

	lines := []string{
		"467..114..",
		"...*......",
	}
	var got [][]string
	for _, line := range lines {
		got = loadInput(got, line)
	}

	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestExtractNumberStructs(t *testing.T) {
	line := []string{"4", "6", "7", ".", ".", "1", "1", "4", ".", "."}
	want := []Number{
		{
			StartIndex: 0,
			EndIndex:   2,
			Value:      467,
		},
		{
			StartIndex: 5,
			EndIndex:   7,
			Value:      114,
		},
	}
	got := extractNumbersFromLine(line)

	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestExtractNumberStructsTwoDigitNumbers(t *testing.T) {
	line := []string{"4", "6", ".", ".", "1", "1", "4", ".", "."}
	want := []Number{
		{
			StartIndex: 0,
			EndIndex:   1,
			Value:      46,
		},
		{
			StartIndex: 4,
			EndIndex:   6,
			Value:      114,
		},
	}
	got := extractNumbersFromLine(line)

	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestExtractNumberStructsSingleDigitNumbers(t *testing.T) {
	line := []string{"4", ".", ".", "1", "1", "4", ".", "."}
	want := []Number{
		{
			StartIndex: 0,
			EndIndex:   0,
			Value:      4,
		},
		{
			StartIndex: 3,
			EndIndex:   5,
			Value:      114,
		},
	}
	got := extractNumbersFromLine(line)

	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

// Basic plan, read the input into a slice of slices
// iterate over each character and for each number, check the adjacent spaces until a symbol is found
// if a symbol is found, concatenate the number characters and convert to an integer
// sum them together

// adjacent spaces will be i+1 and i-1 on the current slice, as well as i-1, i and i+1 on the above and below slices

// should we grab the

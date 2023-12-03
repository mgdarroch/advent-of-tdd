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

func TestMapNumbersToLines(t *testing.T) {
	input := [][]string{
		{"4", "6", "7", ".", ".", "1", "1", "4", ".", "."},
	}
	want := map[int][]Number{
		0: {
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
		},
	}
	got := mapNumbersToLines(input)
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestExtractValidPartsFromTheFirstLine(t *testing.T) {
	input := [][]string{
		{"4", "6", "7", ".", ".", "1", "1", "4", ".", "."},
		{".", ".", ".", "*", ".", ".", ".", ".", ".", "."},
	}
	numberMap := map[int][]Number{
		0: {
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
		},
		1: {},
	}

	want := []Number{
		{
			StartIndex: 0,
			EndIndex:   2,
			Value:      467,
		},
	}

	got := extractValidPartNumbers(input, numberMap)

	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestExtractValidPartsFromTheLastLine(t *testing.T) {
	input := [][]string{
		{".", ".", ".", "*", ".", ".", ".", ".", ".", "."},
		{"4", "6", "7", ".", ".", "1", "1", "4", ".", "."},
	}
	numberMap := map[int][]Number{
		0: {},
		1: {
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
		},
	}

	want := []Number{
		{
			StartIndex: 0,
			EndIndex:   2,
			Value:      467,
		},
	}

	got := extractValidPartNumbers(input, numberMap)

	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestExtractValidPartsFromTheLineWithPreviousAndNext(t *testing.T) {
	input := [][]string{
		{"4", "6", "7", ".", ".", "1", "1", "4", ".", "."},
		{".", ".", ".", "*", ".", ".", ".", ".", ".", "."},
		{".", ".", "3", "5", ".", ".", "6", "3", "3", "."},
		{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
	}
	numberMap := map[int][]Number{
		0: {
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
		},
		1: {},
		2: {
			{
				StartIndex: 2,
				EndIndex:   3,
				Value:      35,
			},
			{
				StartIndex: 6,
				EndIndex:   8,
				Value:      633,
			},
		},
	}

	want := []Number{
		{
			StartIndex: 0,
			EndIndex:   2,
			Value:      467,
		},
		{
			StartIndex: 2,
			EndIndex:   3,
			Value:      35,
		},
		{
			StartIndex: 6,
			EndIndex:   8,
			Value:      633,
		},
	}

	got := extractValidPartNumbers(input, numberMap)

	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestSumValidPartNumbers(t *testing.T) {
	input := []Number{
		{
			StartIndex: 0,
			EndIndex:   2,
			Value:      467,
		},
		{
			StartIndex: 2,
			EndIndex:   3,
			Value:      35,
		},
		{
			StartIndex: 6,
			EndIndex:   8,
			Value:      633,
		},
	}

	want := 1135
	got := 0
	got = sumValidPartNumbers(got, input)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

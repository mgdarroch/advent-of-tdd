package day4

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestBuildNumberMapFromLine(t *testing.T) {
	line := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	want := map[int]int{
		41: 1,
		48: 2,
		83: 2,
		86: 2,
		17: 2,
		6:  1,
		31: 1,
		9:  1,
		53: 1,
	}
	got := buildNumberMap(line)

	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestGetCardValueFromNumberMap(t *testing.T) {
	numberMap := map[int]int{
		41: 1,
		48: 2,
		83: 2,
		86: 2,
		17: 2,
		6:  1,
		31: 1,
		9:  1,
		53: 1,
	}
	want := 8
	got := getCardValue(numberMap)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSolveWithExampleInput(t *testing.T) {
	filePath := "resources/input_test.txt"
	want := 13
	got := Solve(filePath)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

// Part 2

func TestBuildCardStruct(t *testing.T) {
	line := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	numberMap := map[int]int{
		41: 1,
		48: 2,
		83: 2,
		86: 2,
		17: 2,
		6:  1,
		31: 1,
		9:  1,
		53: 1,
	}
	want := Card{
		CardNumber:     1,
		WinningNumbers: 4,
	}
	got := buildCardFromLine(line)

	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

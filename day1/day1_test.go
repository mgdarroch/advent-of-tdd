package day1_test

import (
	"aoctdd/day1"
	"testing"
)

func TestParseLineFindsTheFirstAndLastNumbersInALineWithTwoNumbers(t *testing.T) {
	t.Parallel()
	str := "1abc2"
	want := 12
	got := day1.ParseLine(str)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestParseLineFindsTheFirstAndLastNumbersInALineWithTwoNumbersWithCharactersAtTheStart(t *testing.T) {
	t.Parallel()
	str := "pqr3stu8vwx"
	want := 38
	got := day1.ParseLine(str)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestParseLineFindsTheFirstAndLastNumbersWithMultipleNumbers(t *testing.T) {
	t.Parallel()
	str := "a1b2c3d4e5f"
	want := 15
	got := day1.ParseLine(str)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestParseLineFindsTheFirstAndLastNumbersWithASingleNumber(t *testing.T) {
	t.Parallel()
	str := "treb7uchet"
	want := 77
	got := day1.ParseLine(str)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSumLinesWillAddAllResults(t *testing.T) {
	t.Parallel()
	want := 142
	got := 0
	lines := [4]string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	for _, line := range lines {
		got = day1.SumLines(got, line)
	}
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

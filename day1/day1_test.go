package day1_test

import (
	"aoctdd/day1"
	"testing"
)

//1abc2
//pqr3stu8vwx
//a1b2c3d4e5f
//treb7uchet

func TestFindsTheFirstAndLastNumbersInALineWithTwoNumbers(t *testing.T) {
	t.Parallel()
	str := "1abc2"
	want := 12
	got := day1.ParseLine(str)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestFindsTheFirstAndLastNumbersInALineWithTwoNumbersWithCharactersAtTheStart(t *testing.T) {
	t.Parallel()
	str := "pqr3stu8vwx"
	want := 38
	got := day1.ParseLine(str)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestFindsTheFirstAndLastNumbersWithMultipleNumbers(t *testing.T) {
	t.Parallel()
	str := "a1b2c3d4e5f"
	want := 15
	got := day1.ParseLine(str)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestFindsTheFirstAndLastNumbersWithASingleNumber(t *testing.T) {
	t.Parallel()
	str := "treb7uchet"
	want := 77
	got := day1.ParseLine(str)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

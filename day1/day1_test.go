package day1

import (
	"testing"
)

// PART 1

func TestParseLineFindsTheFirstAndLastNumbersInALineWithTwoNumbers(t *testing.T) {
	t.Parallel()
	str := "1abc2"
	want := 12
	got := parseLine(str)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestParseLineFindsTheFirstAndLastNumbersInALineWithTwoNumbersWithCharactersAtTheStart(t *testing.T) {
	t.Parallel()
	str := "pqr3stu8vwx"
	want := 38
	got := parseLine(str)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestParseLineFindsTheFirstAndLastNumbersWithMultipleNumbers(t *testing.T) {
	t.Parallel()
	str := "a1b2c3d4e5f"
	want := 15
	got := parseLine(str)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestParseLineFindsTheFirstAndLastNumbersWithASingleNumber(t *testing.T) {
	t.Parallel()
	str := "treb7uchet"
	want := 77
	got := parseLine(str)
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
		got = SumLines(got, line)
	}
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

// PART 2

func TestParseDigitStringCanFindValidDigitsWhenThereIsOnlyOne(t *testing.T) {
	t.Parallel()
	line := "two"
	want := 2
	got, _ := parseDigitString(line, 0)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestParseDigitReturnsAnErrorWhereThereIsNoValidDigitStrings(t *testing.T) {
	t.Parallel()
	line := "fail"
	got, err := parseDigitString(line, 0)
	if err == nil {
		t.Errorf("expected no valid strings error, got %d", got)
	}
}

func TestParseLineCanFindValidDigitStringsInLine(t *testing.T) {
	t.Parallel()
	line := "two1nine"
	want := 29
	got := parseLine(line)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestParseLineCanFindValidDigitStringsInAllTestLines(t *testing.T) {
	t.Parallel()
	lines := [7]string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	want := 281
	got := 0
	for _, line := range lines {
		got = SumLines(got, line)
	}
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

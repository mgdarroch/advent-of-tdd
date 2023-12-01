package day1

import (
	"errors"
	"strconv"
	"strings"
)

type DigitString struct {
	Name    string
	Size    int
	Integer int
}

func ParseLine(input string) int {
	var first, last int
	firstSet, lastSet := false, false
	for _, c := range strings.Split(input, "") {
		i, err := strconv.Atoi(c)
		if err != nil {
			continue
		}
		if !firstSet {
			first = i
			firstSet = true
			continue
		}
		last = i
		lastSet = true
	}
	if lastSet == false {
		last = first
	}
	return first*10 + last
}

func SumLines(sum int, input string) int {
	result := sum
	result += ParseLine(input)
	return result
}

func ParseDigitString(input string, start int) (int, error) {
	firstCharMap := map[string][]DigitString{
		"o": {{"one", 3, 1}},
		"t": {{"two", 3, 2}, {"three", 5, 3}},
		"f": {{"four", 4, 4}, {"five", 4, 5}},
		"s": {{"six", 3, 6}, {"seven", 5, 7}},
		"e": {{"eight", 5, 8}},
		"n": {{"nine", 4, 9}},
	}
	options := firstCharMap[input[start:start+1]]
	for _, option := range options {
		strLen := len(input)
		if option.Size > strLen {
			continue
		}
		if option.Size > strLen-start {
			continue
		}
		str := input[start : start+option.Size]
		if option.Name == str {
			return option.Integer, nil
		}
	}
	return 0, errors.New("no valid digits")
}

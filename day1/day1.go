package day1

import (
	"errors"
	"strconv"
	"strings"
)

type DigitString struct {
	Name    string
	Integer int
}

func ParseLine(input string) int {
	var first, last int
	firstSet, lastSet := false, false
	for i, c := range strings.Split(input, "") {
		num, err := strconv.Atoi(c)
		if err != nil {
			digitFromString, err := ParseDigitString(input, i)
			if err != nil {
				continue
			}
			if !firstSet {
				first = digitFromString
				firstSet = true
				continue
			} else {
				last = digitFromString
				lastSet = true
				continue
			}
		}
		if !firstSet {
			first = num
			firstSet = true
			continue
		}
		last = num
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
		"o": {{"one", 1}},
		"t": {{"two", 2}, {"three", 3}},
		"f": {{"four", 4}, {"five", 5}},
		"s": {{"six", 6}, {"seven", 7}},
		"e": {{"eight", 8}},
		"n": {{"nine", 9}},
	}
	options := firstCharMap[input[start:start+1]]
	for _, option := range options {
		strLen := len(input)
		if len(option.Name) > strLen {
			continue
		}
		if len(option.Name) > strLen-start {
			continue
		}
		str := input[start : start+len(option.Name)]
		if option.Name == str {
			return option.Integer, nil
		}
	}
	return 0, errors.New("no valid digit strings")
}

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

func parseLine(input string, digitStrings bool) int {
	var first, last int
	firstSet, lastSet := false, false
	for i, c := range strings.Split(input, "") {
		num, err := strconv.Atoi(c)
		if err != nil {
			if !digitStrings {
				continue
			}
			digitFromString, err := parseDigitString(input, i)
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

func Solve(sum int, input string, digitStrings bool) int {
	result := sum
	result += parseLine(input, digitStrings)
	return result
}

func parseDigitString(input string, start int) (int, error) {
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
	return 0, errors.New("no valid digit strings")
}

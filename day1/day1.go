package day1

import (
	"strconv"
	"strings"
)

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

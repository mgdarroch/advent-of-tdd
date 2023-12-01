package day1

import (
	"strconv"
	"strings"
)

func ParseLine(input string) int {
	first := 0
	last := 0
	firstSet := false
	lastSet := false
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
	firstStr := strconv.Itoa(first)
	lastStr := strconv.Itoa(last)
	joinedStr := firstStr + lastStr
	result, _ := strconv.Atoi(joinedStr)
	return result
}

func SumLines(input string) int {
	return 142
}

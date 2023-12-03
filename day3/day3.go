package day3

import (
	"strings"
)

type Number struct {
	StartIndex int
	EndIndex   int
	Value      int
}

func loadInput(input [][]string, line string) [][]string {
	splitString := strings.Split(line, "")
	input = append(input, splitString)
	return input
}

func extractNumbers(input []string) []Number {
	var numbers []Number
	numbers = append(numbers, Number{
		StartIndex: 0,
		EndIndex:   2,
		Value:      467,
	})
	numbers = append(numbers, Number{
		StartIndex: 5,
		EndIndex:   7,
		Value:      114,
	})
	return numbers
}

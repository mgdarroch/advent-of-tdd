package day3

import (
	"strconv"
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

func extractNumbersFromLine(input []string) []Number {
	var numbers []Number
	for i := 0; i < len(input); i++ {
		number := Number{
			StartIndex: 0,
			EndIndex:   0,
			Value:      0,
		}
		numString := ""
		_, err := strconv.Atoi(input[i])
		if err != nil {
			continue
		}
		number.StartIndex = i
		numString += input[i]
		for j := 1; j < 3; j++ {
			if i+j > len(input) {
				break
			}
			_, err := strconv.Atoi(input[i+j])
			if err == nil {
				numString += input[i+j]
				number.EndIndex = i + j
			}
		}
		number.Value, _ = strconv.Atoi(numString)
		i = i + len(numString)
		numbers = append(numbers, number)
	}
	return numbers
}

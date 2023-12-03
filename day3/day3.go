package day3

import "strings"

func loadInput(input [][]string, line string) [][]string {
	splitString := strings.Split(line, "")
	input = append(input, splitString)
	return input
}

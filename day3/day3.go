package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func mapNumbersToLines(input [][]string) map[int][]Number {
	lineToNumbersMap := map[int][]Number{}
	for i, v := range input {
		lineToNumbersMap[i] = extractNumbersFromLine(v)
	}
	return lineToNumbersMap
}

func extractValidPartNumbers(input [][]string, numberMap map[int][]Number) []Number {
	var validParts []Number
	for i := 0; i < len(input); i++ {
		if i == 0 {
			// first line, only check the next line
			for _, number := range numberMap[i] {
				start := number.StartIndex - 1
				end := number.EndIndex + 1
				if start < 0 {
					start = start + 1
				}
				if end >= len(input[i+1]) {
					end = end - 1
				}
				for first := start; first <= end; first++ {
					if input[i+1][first] != "." {
						_, err := strconv.Atoi(input[i+1][first])
						if err != nil {
							validParts = append(validParts, number)
						}
					}
				}
			}
			continue
		}
		if i == len(input)-1 {
			// last line, only check the previous line
			for _, number := range numberMap[i] {
				start := number.StartIndex - 1
				end := number.EndIndex + 1
				if start < 0 {
					start = start + 1
				}
				if end >= len(input[i-1]) {
					end = end - 1
				}
				for first := start; first <= end; first++ {
					if input[i-1][first] != "." {
						_, err := strconv.Atoi(input[i-1][first])
						if err != nil {
							validParts = append(validParts, number)
						}
					}
				}
			}
			continue
		}
		// check both previous and next line
		for _, number := range numberMap[i] {
			start := number.StartIndex - 1
			end := number.EndIndex + 1
			if start < 0 {
				start = start + 1
			}
			if end >= len(input[i-1]) {
				end = end - 1
			}
			for first := start; first <= end; first++ {
				if input[i-1][first] != "." {
					_, err := strconv.Atoi(input[i-1][first])
					if err != nil {
						validParts = append(validParts, number)
					}
				}
				if input[i+1][first] != "." {
					_, err := strconv.Atoi(input[i+1][first])
					if err != nil {
						validParts = append(validParts, number)
					}
				}
			}
		}
	}
	return validParts
}

func Solve(pathToInput string) {
	f, err := os.Open(pathToInput)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	var input [][]string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		input = loadInput(input, line)
	}

	numberMap := mapNumbersToLines(input)
	fmt.Println(numberMap[1])
}

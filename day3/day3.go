package day3

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type GearPosition struct {
	SymbolIndex int
	Line        int
}

type Number struct {
	StartIndex int
	EndIndex   int
	Value      int
}

func lessNumber(num1, num2 Number) bool {
	return num1.EndIndex < num2.EndIndex
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
		number.EndIndex = i
		numString += input[i]
		for j := 1; j < 3; j++ {
			if i+j > len(input) {
				break
			}
			_, err := strconv.Atoi(input[i+j])
			if err != nil {
				break
			} else {
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

func mapGearPositions(input [][]string) map[GearPosition][]Number {
	lineToNumbersMap := map[GearPosition][]Number{}
	for i, v := range input {
		for j, c := range v {
			if c == "*" {
				gearPos := GearPosition{
					SymbolIndex: j,
					Line:        i,
				}
				lineToNumbersMap[gearPos] = []Number{}
			}
		}
	}
	return lineToNumbersMap
}

func extractValidPartNumbers(input [][]string, numberMap map[int][]Number, gearPosMap map[GearPosition][]Number) []Number {
	var validParts []Number

	checkAndAppend := func(index int, line int, number Number) {
		if index < 0 || index >= len(input[line]) {
			return
		}
		symbol := input[line][index]
		if symbol != "." {
			_, err := strconv.Atoi(symbol)
			if err != nil {
				if symbol == "*" {
					gearPos := GearPosition{
						SymbolIndex: index,
						Line:        line,
					}
					gearPosMap[gearPos] = append(gearPosMap[gearPos], number)
				}
				validParts = append(validParts, number)
			}
		}
	}

	for i, numbers := range numberMap {
		for _, number := range numbers {
			start := max(number.StartIndex-1, 0)
			end := min(number.EndIndex+1, len(input[i])-1)

			// Check the current line
			checkAndAppend(start, i, number)
			checkAndAppend(end, i, number)

			// Check the previous and next line if applicable
			if i > 0 {
				for j := start; j <= end; j++ {
					checkAndAppend(j, i-1, number)
				}
			}
			if i < len(input)-1 {
				for j := start; j <= end; j++ {
					checkAndAppend(j, i+1, number)
				}
			}
		}
	}

	return validParts
}

func sumValidPartNumbers(sum int, numbers []Number) int {
	newSum := sum
	for _, number := range numbers {
		newSum += number.Value
	}
	return newSum
}

func sumGearRatios(gearPosMap map[GearPosition][]Number) int {
	gearRatioSum := 0
	for _, v := range gearPosMap {
		if len(v) == 2 {
			gearRatioSum += v[0].Value * v[1].Value
		}
	}
	return gearRatioSum
}

func Solve(pathToInput string) (int, int) {
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
	gearPosMap := mapGearPositions(input)

	var validParts []Number
	validParts = extractValidPartNumbers(input, numberMap, gearPosMap)

	gearRatioSum := sumGearRatios(gearPosMap)

	answer := 0
	answer = sumValidPartNumbers(answer, validParts)
	return answer, gearRatioSum
}

package day4

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func buildNumberMap(input string) map[int]int {
	numberMap := map[int]int{}
	trimStr := input[8:]
	foundNum := 0
	for i := 0; i < len(trimStr); i++ {
		num, err := strconv.Atoi(string(trimStr[i]))
		if err == nil {
			foundNum = (foundNum * 10) + num
		} else {
			if foundNum != 0 {
				numberMap[foundNum] = numberMap[foundNum] + 1
				foundNum = 0
			}
		}
	}
	if foundNum != 0 {
		numberMap[foundNum] = numberMap[foundNum] + 1
	}
	return numberMap
}

func getCardValue(input map[int]int) int {
	count := 0
	for _, v := range input {
		if v == 2 && count == 0 {
			count += 1
			continue
		}
		if v == 2 {
			count = count * 2
		}
	}
	return count
}

func Solve(inputFile string) int {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	scanner := bufio.NewScanner(f)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		sum += getCardValue(buildNumberMap(line))
	}
	return sum
}

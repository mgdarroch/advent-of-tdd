package day4

import (
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

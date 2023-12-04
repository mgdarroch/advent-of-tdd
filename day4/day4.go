package day4

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Card struct {
	CardNumber     int
	MatchedNumbers int
}

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

func getCardValue(input map[int]int, sum int, matchedNumbers int) (int, int) {
	count := 0
	matches := 0
	for _, v := range input {
		if v == 2 && count == 0 {
			count += 1
			matches++
			continue
		}
		if v == 2 {
			count = count * 2
			matches++
		}
	}
	return sum + count, matches
}

func buildCardFromLine(cardNumber int, matches int) Card {
	return Card{
		CardNumber:     cardNumber,
		MatchedNumbers: matches,
	}
}

func populateCardMapWithDuplicatesAndTotal(cardMap map[int][]Card) (map[int][]Card, int) {
	return map[int][]Card{
		1: {
			{
				CardNumber:     1,
				MatchedNumbers: 4,
			},
		},
		2: {
			{
				CardNumber:     2,
				MatchedNumbers: 2,
			},
			{
				CardNumber:     2,
				MatchedNumbers: 2,
			},
		},
	}, 3
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
		sum, _ = getCardValue(buildNumberMap(line), sum, 0)
	}
	return sum
}

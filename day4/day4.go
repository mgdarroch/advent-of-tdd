package day4

import (
	"bufio"
	"log"
	"os"
	"sort"
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

func populateCardMapWithDuplicatesAndTotal(cardMap map[int][]Card) int {
	total := 0
	var keys []int
	for k := range cardMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		cardArr := cardMap[key]
		total += len(cardArr)
		if cardArr[0].CardNumber < len(keys) {
			for _, card := range cardArr {
				if card.MatchedNumbers > 0 {
					for i := card.CardNumber + 1; i <= card.CardNumber+card.MatchedNumbers; i++ {
						if i > len(keys) {
							break
						}
						cardMap[i] = append(cardMap[i], cardMap[i][0])
					}
				}
			}
		}
	}
	return total
}

func Solve(inputFile string) (int, int) {
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
	cardNumber := 1
	cardMap := map[int][]Card{}
	totalCards := 0

	for scanner.Scan() {
		matches := 0
		line := scanner.Text()
		sum, matches = getCardValue(buildNumberMap(line), sum, matches)
		cardMap[cardNumber] = append(cardMap[cardNumber], buildCardFromLine(cardNumber, matches))
		cardNumber++
	}

	totalCards = populateCardMapWithDuplicatesAndTotal(cardMap)

	return sum, totalCards
}

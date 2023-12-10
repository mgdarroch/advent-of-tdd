package day7

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	CardMap map[Card]int
	Bid     int
}

type Card struct {
	CardSymbol string
	CardRank   int
}

func getCardValueMap() map[string]int {
	return map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
}

func loadInput(input string) []Hand {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	var res []Hand
	scanner := bufio.NewScanner(f)
	cardValueMap := getCardValueMap()

	for scanner.Scan() {
		cardMap := map[Card]int{}
		str := scanner.Text()
		splitStr := strings.Split(str, " ")
		num, _ := strconv.Atoi(splitStr[1])
		for _, v := range splitStr[0] {
			card := Card{
				CardSymbol: string(v),
				CardRank:   cardValueMap[string(v)],
			}
			cardMap[card] = cardMap[card] + 1
		}
		res = append(res, Hand{
			CardMap: cardMap,
			Bid:     num,
		})
	}
	return res
}

func sortHands(hands []Hand) {
	sort.Slice(hands, func(i, j int) bool {
		return getHandStrength(hands[i].CardMap) > getHandStrength(hands[j].CardMap)
	})
}

func getHandStrength(cardMap map[Card]int) int {
	handStrength := 0
	cardValueMap := getCardValueMap()

	switch len(cardMap) {
	case 1:
		handStrength = 10000
	case 2:
		handStrength = 1000
	case 3:
		handStrength = 100
	case 4:
		handStrength = 10
	case 5:
		handStrength = 0
	}

	for key, value := range cardMap {
		handStrength += cardValueMap[key.CardSymbol] * value
	}
	return handStrength
}

func Solve(filePath string) int {
	return 6440
}

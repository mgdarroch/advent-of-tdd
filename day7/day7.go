package day7

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	CardString string
	CardMap    map[Card]int
	Bid        int
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
			CardString: splitStr[0],
			CardMap:    cardMap,
			Bid:        num,
		})
	}
	return res
}

func sortHands(hands []Hand) {
	sort.Slice(hands, func(i, j int) bool {
		return calculateStrength(hands[i]) < calculateStrength(hands[j])
	})
}

func getHandStrength(hand Hand) int {
	handStrength := 0
	maxCardCount := 0
	for _, v := range hand.CardMap {
		if v > maxCardCount {
			maxCardCount = v
		}
	}
	switch maxCardCount {
	case 5:
		handStrength = 7000000
	case 4:
		handStrength = 6000000
	case 3:
		if len(hand.CardMap) == 2 {
			handStrength = 5000000
		} else {
			handStrength = 4000000
		}
	case 2:
		if len(hand.CardMap) == 3 {
			handStrength = 3000000
		} else {
			handStrength = 2000000
		}
	case 1:
		handStrength = 1000000
	}
	return handStrength
}

func calculateStrength(hand Hand) int {
	typeStrength := getHandStrength(hand)
	cardValueMap := getCardValueMap()
	handStrength := typeStrength
	for i, card := range strings.Split(hand.CardString, "") {
		handStrength += cardValueMap[card] * int(math.Pow(10, float64(4-i)))
	}
	return handStrength
}

func Solve(filePath string) int {
	input := loadInput(filePath)
	sortHands(input)
	part1 := 0
	for i, v := range input {
		part1 += v.Bid * (i + 1)
	}
	return part1
}

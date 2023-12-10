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
	CardString   string
	CardMap      map[Card]int
	Bid          int
	Type         string
	HandStrength int
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
			Type:       "",
		})
	}
	return res
}

func sortHands(hands []Hand) {
	sort.Slice(hands, func(i, j int) bool {
		if calculateStrength(&hands[i]) < calculateStrength(&hands[j]) {
			return true
		}
		return false
	})
}

func getHandStrength(hand *Hand) int {
	var handStrength int
	maxCardCount := 0
	for _, v := range hand.CardMap {
		if v > maxCardCount {
			maxCardCount = v
		}
	}
	switch maxCardCount {
	case 5:
		hand.Type = "Five of a Kind"
		handStrength = 7
	case 4:
		hand.Type = "Four of a Kind"
		handStrength = 6
	case 3:
		if len(hand.CardMap) == 2 {
			hand.Type = "Full House"
			handStrength = 5
		} else {
			hand.Type = "Three of a Kind"
			handStrength = 4
		}
	case 2:
		if len(hand.CardMap) == 3 {
			hand.Type = "Two Pair"
			handStrength = 3
		} else {
			hand.Type = "One Pair"
			handStrength = 2
		}
	case 1:
		hand.Type = "High Card"
		handStrength = 1
	}
	return handStrength
}

func calculateStrength(hand *Hand) int {
	typeStrength := getHandStrength(hand)
	typeStr := strconv.Itoa(typeStrength)
	cardValueMap := getCardValueMap()
	cards := strings.Split(hand.CardString, "")
	for i := 0; i < len(hand.CardString); i++ {
		if cardValueMap[cards[i]] < 10 {
			typeStr += "0"
			typeStr += strconv.Itoa(cardValueMap[cards[i]])
		} else {
			typeStr += strconv.Itoa(cardValueMap[cards[i]])
		}
	}
	res, _ := strconv.Atoi(typeStr)
	hand.HandStrength = res
	return res
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

package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func getHandStrength(hand Hand) int {
	handStrength := 0
	handMap := map[string]int{}
	fmt.Println(handMap)
	fmt.Println(hand.CardMap)
	fmt.Println(handStrength)
	return 440
}

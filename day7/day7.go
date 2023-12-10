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
	CardMap map[string]int
	Bid     int
}

type Card struct {
	CardSymbol string
	CardRank   int
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
	for scanner.Scan() {
		cardMap := map[string]int{}
		str := scanner.Text()
		splitStr := strings.Split(str, " ")
		num, _ := strconv.Atoi(splitStr[1])
		for _, v := range splitStr[0] {
			cardMap[string(v)] = cardMap[string(v)] + 1
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

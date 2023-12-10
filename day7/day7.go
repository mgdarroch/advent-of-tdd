package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Round struct {
	Hand Hand
	Bid  int
}

type Hand struct {
	Cards string
}

type Card struct {
	CardSymbol string
	CardRank   int
}

func loadInput(input string) []Round {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	var res []Round

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		splitStr := strings.Split(str, " ")
		num, _ := strconv.Atoi(splitStr[1])
		res = append(res, Round{
			Hand: Hand{
				Cards: splitStr[0],
			},
			Bid: num,
		})
	}
	return res
}

func getHandStrength(round Round) int {
	handStrength := 0
	handMap := map[string]int{}
	for _, v := range round.Hand.Cards {
		handMap[string(v)] = handMap[string(v)] + 1
	}
	fmt.Println(handStrength)
	return 440
}

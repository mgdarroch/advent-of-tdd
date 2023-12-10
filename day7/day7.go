package day7

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Round struct {
	Hand string
	Bid  int
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
			Hand: splitStr[0],
			Bid:  num,
		})
	}
	return res
}

func getHandStrength(round Round) int {
	return 440
}

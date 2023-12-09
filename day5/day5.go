package day5

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func loadInput(input string) []string {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	var res []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

func parseSeeds(input string) []int {
	seeds := input[7:]
	seedsSplit := strings.Split(seeds, " ")
	var intSeeds []int
	for _, seed := range seedsSplit {
		num, _ := strconv.Atoi(seed)
		intSeeds = append(intSeeds, num)
	}
	return intSeeds
}

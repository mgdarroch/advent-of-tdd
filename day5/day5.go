package day5

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	source      int
	destination int
	transform   int
}

func loadInput(input string) [][]string {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	var res [][]string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		str := scanner.Text()
		res = append(res, strings.Split(str, " "))
	}
	return res
}

func parseSeeds(input []string) []int {
	seeds := input[1:]
	var intSeeds []int
	for _, seed := range seeds {
		num, _ := strconv.Atoi(seed)
		intSeeds = append(intSeeds, num)
	}
	return intSeeds
}

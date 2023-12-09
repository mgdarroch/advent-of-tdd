package day5

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	From      int
	To        int
	Transform int
}

type Map struct {
	Name   string
	Ranges []Range
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
		if str == "" {
			continue
		}
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

func parseMap(lines [][]string) Map {
	res := Map{
		Name:   "",
		Ranges: []Range{},
	}
	for i, line := range lines {
		if i == 0 {
			res.Name = line[0]
			continue
		}
		sourceRange, _ := strconv.Atoi(line[0])
		destRange, _ := strconv.Atoi(line[1])
		transform, _ := strconv.Atoi(line[2])
		res.Ranges = append(res.Ranges, Range{
			From:      sourceRange,
			To:        destRange,
			Transform: transform,
		})
	}

	return res
}

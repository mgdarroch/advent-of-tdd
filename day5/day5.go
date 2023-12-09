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

func extractMaps(lines [][]string) []Map {
	return []Map{
		{
			Name: "seed-to-soil",
			Ranges: []Range{
				{50, 98, 2},
				{52, 50, 48},
			},
		},
		{
			Name: "soil-to-fertilizer",
			Ranges: []Range{
				{0, 15, 37},
				{37, 52, 2},
				{39, 0, 15},
			},
		},
		{
			Name: "fertilizer-to-water",
			Ranges: []Range{
				{49, 53, 8},
				{0, 11, 42},
				{42, 0, 7},
				{57, 7, 4},
			},
		},
		{
			Name: "water-to-light",
			Ranges: []Range{
				{88, 18, 7},
				{18, 25, 70},
			},
		},
		{
			Name: "light-to-temperature",
			Ranges: []Range{
				{45, 77, 23},
				{81, 45, 19},
				{68, 64, 13},
			},
		},
		{
			Name: "temperature-to-humidity",
			Ranges: []Range{
				{0, 69, 1},
				{1, 0, 69},
			},
		},
		{
			Name: "humidity-to-location",
			Ranges: []Range{
				{60, 56, 37},
				{56, 93, 4},
			},
		},
	}
}

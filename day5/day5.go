package day5

import (
	"bufio"
	"log"
	"os"
	"sort"
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
	var ranges []Range
	for i, line := range lines {
		if i == 0 {
			res.Name = line[0]
			continue
		}
		sourceRange, _ := strconv.Atoi(line[0])
		destRange, _ := strconv.Atoi(line[1])
		rangeLength, _ := strconv.Atoi(line[2])
		ranges = append(ranges, Range{
			From:      sourceRange,
			To:        sourceRange + rangeLength,
			Transform: destRange - sourceRange,
		})
	}
	sort.Slice(ranges, func(c1, c2 int) bool {
		return ranges[c1].From < ranges[c2].From
	})
	res.Ranges = ranges
	return res
}

func extractMaps(lines [][]string) []Map {
	var res []Map

	for i := 1; i < len(lines); i++ {
		if strings.Contains(lines[i][1], "map") {
			for j := i + 1; j < len(lines); j++ {
				if strings.Contains(lines[j][1], "map") {
					res = append(res, parseMap(lines[i:j]))
					i = j - 1
					break
				} else if j == len(lines)-1 {
					res = append(res, parseMap(lines[i:]))
					i = j - 1
				}
			}
		}
	}

	return res
}

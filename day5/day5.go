package day5

import (
	"bufio"
	"log"
	"math"
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
	Ranges []Range
}

type ExtractRangeFunc func(ranges []Range, sourceRange int, destRange int, rangeLength int) []Range

func (m Map) get(v int) (int, bool) {
	// binary search for the range
	l := 0
	r := len(m.Ranges) - 1
	for l <= r {
		mid := l + (r-l)/2
		rng := m.Ranges[mid]
		if v >= rng.From && v <= rng.To {
			return v + rng.Transform, true
		}
		if v < rng.From {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return 0, false
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

func parseMap(lines [][]string, rangeFunc ExtractRangeFunc) Map {
	res := Map{
		Ranges: []Range{},
	}
	var ranges []Range
	for i, line := range lines {
		if i == 0 {
			continue
		}
		destRange, _ := strconv.Atoi(line[0])
		sourceRange, _ := strconv.Atoi(line[1])
		rangeLength, _ := strconv.Atoi(line[2])
		ranges = rangeFunc(ranges, sourceRange, destRange, rangeLength)
	}
	sort.Slice(ranges, func(c1, c2 int) bool {
		return ranges[c1].From < ranges[c2].From
	})
	res.Ranges = ranges
	return res
}

func extractRangesPart1(ranges []Range, sourceRange int, destRange int, rangeLength int) []Range {
	ranges = append(ranges, Range{
		From:      sourceRange,
		To:        sourceRange + rangeLength - 1,
		Transform: destRange - sourceRange,
	})
	return ranges
}

func extractMaps(lines [][]string, rangeFunc ExtractRangeFunc) []Map {
	var res []Map

	for i := 1; i < len(lines); i++ {
		if strings.Contains(lines[i][1], "map") {
			for j := i + 1; j < len(lines); j++ {
				if strings.Contains(lines[j][1], "map") {
					res = append(res, parseMap(lines[i:j], rangeFunc))
					i = j - 1
					break
				} else if j == len(lines)-1 {
					res = append(res, parseMap(lines[i:], rangeFunc))
					i = j - 1
				}
			}
		}
	}

	return res
}

func Solve(input string) (int, int) {
	loadedInput := loadInput(input)
	seeds := parseSeeds(loadedInput[0])
	maps := extractMaps(loadedInput, extractRangesPart1)

	lowest := math.MaxInt
	for _, v := range seeds {
		for _, m := range maps {
			if dst, contains := m.get(v); contains {
				v = dst
			}
		}
		lowest = min(lowest, v)
	}

	seedsWithRange := parseSeedsWithRange(loadedInput[0])
	mapsPart2 := extractMaps(loadedInput, extractRangesPart2)

	for loc := 0; loc < math.MaxInt; loc++ {
		v := transformRev(loc, mapsPart2)
		if isWithinSeedRange(v, seedsWithRange) {
			return lowest, loc
		}
	}

	return lowest, -1
}

func parseSeedsWithRange(input []string) [][2]int {
	seeds := input[1:]
	count := 1
	num1 := 0
	num2 := 0
	var res [][2]int
	for _, seed := range seeds {
		if count == 1 {
			num1, _ = strconv.Atoi(seed)
			count++
			continue
		}
		if count == 2 {
			num2, _ = strconv.Atoi(seed)
			count++
		}
		if count == 3 {
			res = append(res, [2]int{num1, num2})
			count = 1
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	return res
}

func extractRangesPart2(ranges []Range, sourceRange int, destRange int, rangeLength int) []Range {
	ranges = append(ranges, Range{
		From:      destRange,
		To:        destRange + rangeLength - 1,
		Transform: sourceRange - destRange,
	})
	return ranges
}

func transformRev(v int, maps []Map) int {
	for i := len(maps) - 1; i >= 0; i-- {
		m := maps[i]
		if dst, contains := m.get(v); contains {
			v = dst
		}
	}
	return v
}

func isWithinSeedRange(n int, seeds [][2]int) bool {
	l := 0
	r := len(seeds) - 1
	for l <= r {
		mid := l + (r-l)/2
		rng := seeds[mid]
		if n < rng[0] {
			r = mid - 1
		} else if n > rng[0]+rng[1]-1 {
			l = mid + 1
		} else {
			return true
		}
	}

	return false
}

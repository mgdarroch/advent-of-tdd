package day6

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func loadInput(input string) []Race {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	var races []Race
	var lines []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	timeNumbers := strings.Fields(lines[0][11:])
	distanceNumbers := strings.Fields(lines[1][11:])

	for i := 0; i < len(timeNumbers); i++ {
		timeNum, _ := strconv.Atoi(timeNumbers[i])
		distanceNum, _ := strconv.Atoi(distanceNumbers[i])
		races = append(races, Race{
			Time:     timeNum,
			Distance: distanceNum,
		})
	}
	return races
}

func checkWinScenarios(race Race) int {
	winScenarios := 0
	for i := 1; i < race.Time; i++ {
		distance := i * (race.Time - i)
		if distance > race.Distance {
			winScenarios++
		}
	}
	return winScenarios
}

func Solve(filePath string) int {
	input := loadInput(filePath)
	part1 := 1
	for _, race := range input {
		part1 = part1 * checkWinScenarios(race)
	}
	return part1
}

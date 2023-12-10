package main

import (
	"aoctdd/day1"
	"aoctdd/day2"
	"aoctdd/day3"
	"aoctdd/day4"
	"aoctdd/day5"
	"aoctdd/day6"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	solutionDay1()
	solutionDay2()
	solutionDay3()
	solutionDay4()
	solutionDay5()
	solutionDay6()
}

func solutionDay1() {
	f, err := os.Open("day1/resources/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	scanner := bufio.NewScanner(f)
	part1Answer := 0
	part2Answer := 0
	for scanner.Scan() {
		line := scanner.Text()
		part1Answer = day1.Solve(part1Answer, line, false)
		part2Answer = day1.Solve(part2Answer, line, true)
	}
	fmt.Printf("Day 1 => Part 1: %d Part 2: %d\n", part1Answer, part2Answer)
}

func solutionDay2() {
	f, err := os.Open("day2/resources/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	scanner := bufio.NewScanner(f)
	idSum := 0
	powerSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		idSum, powerSum = day2.Solve(line, idSum, powerSum)
	}
	fmt.Printf("Day 2 => Part 1: %d, Part 2: %d\n", idSum, powerSum)
}

func solutionDay3() {
	filePath := "day3/resources/input.txt"
	part1, part2 := day3.Solve(filePath)
	fmt.Printf("Day 3 => Part 1: %d Part 2: %d\n", part1, part2)
}

func solutionDay4() {
	filePath := "day4/resources/input.txt"
	part1, part2 := day4.Solve(filePath)
	fmt.Printf("Day 4 => Part 1: %d, Part 2: %d\n", part1, part2)
}

func solutionDay5() {
	filePath := "day5/resources/input.txt"
	part1, part2 := day5.Solve(filePath)
	fmt.Printf("Day 5 => Part 1: %d, Part 2: %d\n", part1, part2)
}

func solutionDay6() {
	filePath := "day6/resources/input.txt"
	part1, part2 := day6.Solve(filePath)
	fmt.Printf("Day 6 => Part 1: %d, Part 2: %d\n", part1, part2)
}

package main

import (
	"aoctdd/day1"
	"aoctdd/day2"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	solutionDay1()
	solutionDay2()
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
	fmt.Printf("Day 2 => Part 1: %d, Part 2: %d", idSum, powerSum)
}

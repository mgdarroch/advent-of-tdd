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
	defer f.Close()
	scanner := bufio.NewScanner(f)
	answer := 0
	for scanner.Scan() {
		line := scanner.Text()
		answer = day1.SolveDay1(answer, line)
	}
	fmt.Println(answer)
}

func solutionDay2() {
	f, err := os.Open("day2/resources/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	idSum := 0
	powerSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		idSum, powerSum = day2.ValidateGame(line, idSum, powerSum)
	}
	fmt.Printf("ID: %d, Power: %d", idSum, powerSum)
}

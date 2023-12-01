package main

import (
	"aoctdd/day1"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	solutionDay1()
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
		answer = day1.SumLines(answer, line)
	}
	fmt.Println(answer)
}

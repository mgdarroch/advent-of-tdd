package day2_test

import (
	"aoctdd/day2"
	"testing"
)

/*

Test Input

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

Find the games that can contain only 12 red cubes, 13 green cubes, and 14 blue cubes
Game 3 is impossible as at one point 20 red cubes are pulled out.
Game 4 is impossible as at one point 15 blue cubes are pulled out.

So we add up 1, 2, 5 and get the answer of 8.
*/

func TestParseLineShouldReturnTheIdIfGameIsPossible(t *testing.T) {
	t.Parallel()
	validGame := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	want := 1
	got, _ := day2.ParseLine(validGame)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestParseLineShouldReturnErrIfTheGameIsImpossible(t *testing.T) {
	t.Parallel()
	invalidGame := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	got, err := day2.ParseLine(invalidGame)
	if err == nil {
		t.Errorf("wanted an invalid game error, got %d", got)
	}
}

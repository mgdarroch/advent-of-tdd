package day2_test

import (
	"aoctdd/day2"
	"github.com/google/go-cmp/cmp"
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
	games := []day2.Game{
		{
			RedCount:   4,
			BlueCount:  3,
			GreenCount: 0,
		},
		{
			RedCount:   1,
			BlueCount:  6,
			GreenCount: 2,
		},
		{
			RedCount:   0,
			BlueCount:  0,
			GreenCount: 2,
		},
	}
	want := day2.GameInfo{
		Id:       1,
		GameData: games,
	}
	got := day2.ParseLine(validGame)
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestParseLineShouldReturnDoubleDigitIdIfGameIsPossible(t *testing.T) {
	t.Parallel()
	validGame := "Game 10: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	games := []day2.Game{
		{
			RedCount:   4,
			BlueCount:  3,
			GreenCount: 0,
		},
		{
			RedCount:   1,
			BlueCount:  6,
			GreenCount: 2,
		},
		{
			RedCount:   0,
			BlueCount:  0,
			GreenCount: 2,
		},
	}
	want := day2.GameInfo{
		Id:       10,
		GameData: games,
	}
	got := day2.ParseLine(validGame)
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestParseLineShouldReturnTripleDigitIdIfGameIsPossible(t *testing.T) {
	t.Parallel()
	validGame := "Game 100: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	games := []day2.Game{
		{
			RedCount:   4,
			BlueCount:  3,
			GreenCount: 0,
		},
		{
			RedCount:   1,
			BlueCount:  6,
			GreenCount: 2,
		},
		{
			RedCount:   0,
			BlueCount:  0,
			GreenCount: 2,
		},
	}
	want := day2.GameInfo{
		Id:       100,
		GameData: games,
	}
	got := day2.ParseLine(validGame)
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestValidateGameShouldReturnIncrementedSumIfGameIsValid(t *testing.T) {
	t.Parallel()
	want := 1
	line := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	gotSum, _ := day2.ValidateGame(0, line)
	if want != gotSum {
		t.Errorf("want %d, got %d", want, gotSum)
	}
}

func TestValidateGameShouldReturnSameSumIfGameIsInvalid(t *testing.T) {
	t.Parallel()
	want := 0
	line := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	got, _ := day2.ValidateGame(0, line)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestValidateGameShouldReturnSameSumIfGameIsInvalidMoreThanThreeSessions(t *testing.T) {
	t.Parallel()
	want := 0
	line := "Game 11: 2 red, 18 blue, 5 green; 4 green, 12 blue, 9 red; 6 red, 4 green, 5 blue; 8 red, 16 blue, 2 green; 1 green, 18 blue, 13 red; 13 blue, 9 red"
	got, _ := day2.ValidateGame(0, line)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestValidateGameCalledOverExampleInputReturns8(t *testing.T) {
	t.Parallel()
	lines := [5]string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	want := 8
	got := 0
	for _, line := range lines {
		got, _ = day2.ValidateGame(got, line)
	}
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestValidateGameReturnsTheSumOfThePowerOfMinimumCubesTotal48(t *testing.T) {
	t.Parallel()
	wantPower := 48
	line := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	_, gotPower := day2.ValidateGame(0, line)
	if wantPower != gotPower {
		t.Errorf("want %d, got %d", wantPower, gotPower)
	}
}

func TestValidateGameReturnsTheSumOfThePowerOfMinimumCubesTotal1560(t *testing.T) {
	t.Parallel()
	wantPower := 1560
	line := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	_, gotPower := day2.ValidateGame(0, line)
	if wantPower != gotPower {
		t.Errorf("want %d, got %d", wantPower, gotPower)
	}
}

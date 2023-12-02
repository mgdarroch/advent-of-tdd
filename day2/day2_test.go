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
	got, _ := day2.ParseLine(validGame)
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
	got, _ := day2.ParseLine(validGame)
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestParseLineShouldReturnErrIfTheGameIsImpossible(t *testing.T) {
	t.Parallel()
	invalidGame := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	got, err := day2.ParseLine(invalidGame)
	if err == nil {
		t.Errorf("wanted an invalid game error, got %q", got)
	}
}

func TestValidateGameShouldReturnIncrementedSumIfGameIsValid(t *testing.T) {
	t.Parallel()
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
	validGame := day2.GameInfo{
		Id:       10,
		GameData: games,
	}
	want := 10
	got := day2.ValidateGame(0, validGame)
	if want != got {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestValidateGameShouldReturnSameSumIfGameIsInvalid(t *testing.T) {
	t.Parallel()
	games := []day2.Game{
		{
			RedCount:   20,
			BlueCount:  6,
			GreenCount: 8,
		},
		{
			RedCount:   4,
			BlueCount:  5,
			GreenCount: 13,
		},
		{
			RedCount:   1,
			BlueCount:  0,
			GreenCount: 5,
		},
	}
	invalidGame := day2.GameInfo{
		Id:       10,
		GameData: games,
	}
	want := 0
	got := day2.ValidateGame(0, invalidGame)
	if want != got {
		t.Errorf("expected %d, got %d", want, got)
	}
}

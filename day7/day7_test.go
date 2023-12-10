package day7

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLoadInputToRoundStruct(t *testing.T) {
	input := "resources/input_test.txt"
	want := []Hand{
		{
			CardMap: map[string]int{
				"3": 2,
				"2": 1,
				"T": 1,
				"K": 1,
			},
			Bid: 765,
		},
		{
			CardMap: map[string]int{
				"T": 1,
				"5": 3,
				"J": 1,
			},
			Bid: 684,
		},
		{
			CardMap: map[string]int{
				"K": 2,
				"7": 2,
				"6": 1,
			},
			Bid: 28,
		},
		{
			CardMap: map[string]int{
				"K": 1,
				"T": 2,
				"J": 2,
			},
			Bid: 220,
		},
		{
			CardMap: map[string]int{
				"Q": 3,
				"J": 1,
				"A": 1,
			},
			Bid: 483,
		},
	}
	got := loadInput(input)
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestGetHandStrength(t *testing.T) {
	input := Hand{
		CardMap: map[string]int{
			"K": 1,
			"T": 2,
			"J": 2,
		},
		Bid: 220,
	}
	want := 440
	got := getHandStrength(input)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

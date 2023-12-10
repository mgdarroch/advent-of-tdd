package day7

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLoadInputToRoundStruct(t *testing.T) {
	input := "resources/input_test.txt"
	want := []Round{
		{
			Hand: Hand{Cards: "32T3K"},
			Bid:  765,
		},
		{
			Hand: Hand{Cards: "T55J5"},
			Bid:  684,
		},
		{
			Hand: Hand{Cards: "KK677"},
			Bid:  28,
		},
		{
			Hand: Hand{Cards: "KTJJT"},
			Bid:  220,
		},
		{
			Hand: Hand{Cards: "QQQJA"},
			Bid:  483,
		},
	}
	got := loadInput(input)
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestGetHandStrength(t *testing.T) {
	input := Round{
		Hand: Hand{Cards: "KTJJT"},
		Bid:  220,
	}
	want := 440
	got := getHandStrength(input)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

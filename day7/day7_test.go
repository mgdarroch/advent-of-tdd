package day7

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLoadInputToRoundStruct(t *testing.T) {
	input := "resources/input_test.txt"
	want := []Round{
		{
			Hand: "32T3K",
			Bid: 765,
		},
		{
			Hand: "T55J5",
			Bid: 684,
		},
		{
			Hand: "KK677",
			Bid: 28
		},
		{
			Hand: "KTJJT",
			Bid: 220
		},
		{
			Hand: "QQQJA",
			Bid: 483
		},
	}
	got := loadInput(input)
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

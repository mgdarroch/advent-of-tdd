package day7

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLoadInputToRoundStruct(t *testing.T) {
	input := "resources/input_test.txt"
	want := []Hand{
		{
			Cards: "32T3K",
			Bid:   765,
		},
		{
			Cards: "T55J5",
			Bid:   684,
		},
		{
			Cards: "KK677",
			Bid:   28,
		},
		{
			Cards: "KTJJT",
			Bid:   220,
		},
		{
			Cards: "QQQJA",
			Bid:   483,
		},
	}
	got := loadInput(input)
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestGetHandStrength(t *testing.T) {
	input := Hand{
		Cards: "KTJJT",
		Bid:   220,
	}
	want := 440
	got := getHandStrength(input)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

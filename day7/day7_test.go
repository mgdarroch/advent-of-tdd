package day7

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLoadInputToHandStructArray(t *testing.T) {
	input := "resources/input_test.txt"
	want := []Hand{
		{
			CardMap: map[Card]int{
				Card{CardSymbol: "3", CardRank: 3}:  2,
				Card{CardSymbol: "2", CardRank: 2}:  1,
				Card{CardSymbol: "T", CardRank: 10}: 1,
				Card{CardSymbol: "K", CardRank: 13}: 1,
			},
			Bid: 765,
		},
		{
			CardMap: map[Card]int{
				Card{CardSymbol: "T", CardRank: 10}: 1,
				Card{CardSymbol: "5", CardRank: 5}:  3,
				Card{CardSymbol: "J", CardRank: 11}: 1,
			},
			Bid: 684,
		},
		{
			CardMap: map[Card]int{
				Card{CardSymbol: "K", CardRank: 13}: 2,
				Card{CardSymbol: "7", CardRank: 7}:  2,
				Card{CardSymbol: "6", CardRank: 6}:  1,
			},
			Bid: 28,
		},
		{
			CardMap: map[Card]int{
				Card{CardSymbol: "K", CardRank: 13}: 1,
				Card{CardSymbol: "T", CardRank: 10}: 2,
				Card{CardSymbol: "J", CardRank: 11}: 2,
			},
			Bid: 220,
		},
		{
			CardMap: map[Card]int{
				Card{CardSymbol: "Q", CardRank: 12}: 3,
				Card{CardSymbol: "J", CardRank: 11}: 1,
				Card{CardSymbol: "A", CardRank: 14}: 1,
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
		CardMap: map[Card]int{
			Card{CardSymbol: "K", CardRank: 13}: 1,
			Card{CardSymbol: "T", CardRank: 10}: 2,
			Card{CardSymbol: "J", CardRank: 11}: 2,
		},
		Bid: 220,
	}
	want := 440
	got := getHandStrength(input)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

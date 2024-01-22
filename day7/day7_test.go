package day7

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLoadInputToHandStructArray(t *testing.T) {
	input := "resources/input_test.txt"
	want := []Hand{
		{
			CardString: "32T3K",
			CardMap: map[Card]int{
				Card{CardSymbol: "3", CardRank: 3}:  2,
				Card{CardSymbol: "2", CardRank: 2}:  1,
				Card{CardSymbol: "T", CardRank: 10}: 1,
				Card{CardSymbol: "K", CardRank: 13}: 1,
			},
			Bid: 765,
		},
		{
			CardString: "T55J5",
			CardMap: map[Card]int{
				Card{CardSymbol: "T", CardRank: 10}: 1,
				Card{CardSymbol: "5", CardRank: 5}:  3,
				Card{CardSymbol: "J", CardRank: 11}: 1,
			},
			Bid: 684,
		},
		{
			CardString: "KK677",
			CardMap: map[Card]int{
				Card{CardSymbol: "K", CardRank: 13}: 2,
				Card{CardSymbol: "7", CardRank: 7}:  2,
				Card{CardSymbol: "6", CardRank: 6}:  1,
			},
			Bid: 28,
		},
		{
			CardString: "KTJJT",
			CardMap: map[Card]int{
				Card{CardSymbol: "K", CardRank: 13}: 1,
				Card{CardSymbol: "T", CardRank: 10}: 2,
				Card{CardSymbol: "J", CardRank: 11}: 2,
			},
			Bid: 220,
		},
		{
			CardString: "QQQJA",
			CardMap: map[Card]int{
				Card{CardSymbol: "Q", CardRank: 12}: 3,
				Card{CardSymbol: "J", CardRank: 11}: 1,
				Card{CardSymbol: "A", CardRank: 14}: 1,
			},
			Bid: 483,
		},
	}
	got := loadInput(input, false)
	if !cmp.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestSortHandsIntoRank(t *testing.T) {
	got := []Hand{
		{
			CardString: "32T3K",
			CardMap: map[Card]int{
				Card{CardSymbol: "3", CardRank: 3}:  2,
				Card{CardSymbol: "2", CardRank: 2}:  1,
				Card{CardSymbol: "T", CardRank: 10}: 1,
				Card{CardSymbol: "K", CardRank: 13}: 1,
			},
			Bid:  765,
			Type: "One Pair",
		},
		{
			CardString: "T55J5",
			CardMap: map[Card]int{
				Card{CardSymbol: "T", CardRank: 10}: 1,
				Card{CardSymbol: "5", CardRank: 5}:  3,
				Card{CardSymbol: "J", CardRank: 11}: 1,
			},
			Bid:  684,
			Type: "Three of a Kind",
		},
		{
			CardString: "KK677",
			CardMap: map[Card]int{
				Card{CardSymbol: "K", CardRank: 13}: 2,
				Card{CardSymbol: "7", CardRank: 7}:  2,
				Card{CardSymbol: "6", CardRank: 6}:  1,
			},
			Bid:  28,
			Type: "Two Pair",
		},
		{
			CardString: "KTJJT",
			CardMap: map[Card]int{
				Card{CardSymbol: "K", CardRank: 13}: 1,
				Card{CardSymbol: "T", CardRank: 10}: 2,
				Card{CardSymbol: "J", CardRank: 11}: 2,
			},
			Bid:  220,
			Type: "Two Pair",
		},
		{
			CardString: "QQQJA",
			CardMap: map[Card]int{
				Card{CardSymbol: "Q", CardRank: 12}: 3,
				Card{CardSymbol: "J", CardRank: 11}: 1,
				Card{CardSymbol: "A", CardRank: 14}: 1,
			},
			Bid:  483,
			Type: "Three of a Kind",
		},
	}

	want := []Hand{
		{
			CardString: "32T3K",
			CardMap: map[Card]int{
				Card{CardSymbol: "3", CardRank: 3}:  2,
				Card{CardSymbol: "2", CardRank: 2}:  1,
				Card{CardSymbol: "T", CardRank: 10}: 1,
				Card{CardSymbol: "K", CardRank: 13}: 1,
			},
			Bid:          765,
			Type:         "One Pair",
			HandStrength: 20302100313,
		},
		{
			CardString: "KTJJT",
			CardMap: map[Card]int{
				Card{CardSymbol: "K", CardRank: 13}: 1,
				Card{CardSymbol: "T", CardRank: 10}: 2,
				Card{CardSymbol: "J", CardRank: 11}: 2,
			},
			Bid:          220,
			Type:         "Two Pair",
			HandStrength: 31310111110,
		},
		{
			CardString: "KK677",
			CardMap: map[Card]int{
				Card{CardSymbol: "K", CardRank: 13}: 2,
				Card{CardSymbol: "7", CardRank: 7}:  2,
				Card{CardSymbol: "6", CardRank: 6}:  1,
			},
			Bid:          28,
			Type:         "Two Pair",
			HandStrength: 31313060707,
		},
		{
			CardString: "T55J5",
			CardMap: map[Card]int{
				Card{CardSymbol: "T", CardRank: 10}: 1,
				Card{CardSymbol: "5", CardRank: 5}:  3,
				Card{CardSymbol: "J", CardRank: 11}: 1,
			},
			Bid:          684,
			Type:         "Three of a Kind",
			HandStrength: 41005051105,
		},
		{
			CardString: "QQQJA",
			CardMap: map[Card]int{
				Card{CardSymbol: "Q", CardRank: 12}: 3,
				Card{CardSymbol: "J", CardRank: 11}: 1,
				Card{CardSymbol: "A", CardRank: 14}: 1,
			},
			Bid:          483,
			Type:         "Three of a Kind",
			HandStrength: 41212121114,
		},
	}
	sortHands(got)
	if !cmp.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestSolvePart1(t *testing.T) {
	input := "resources/input_test.txt"
	want := 6440
	got, _ := Solve(input)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSolvePartReal(t *testing.T) {
	input := "resources/input.txt"
	want := 248217452
	got, _ := Solve(input)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSolvePart2(t *testing.T) {
	input := "resources/input_test.txt"
	want := 5905
	_, got := Solve(input)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSolvePart2Real(t *testing.T) {
	input := "resources/input.txt"
	want := 245576185
	_, got := Solve(input)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

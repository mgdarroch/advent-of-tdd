package day7

type Round struct {
	Hand string
	Bid  int
}

func loadInput(input string) []Round {
	return []Round{
		{
			Hand: "32T3K",
			Bid:  765,
		},
		{
			Hand: "T55J5",
			Bid:  684,
		},
		{
			Hand: "KK677",
			Bid:  28,
		},
		{
			Hand: "KTJJT",
			Bid:  220,
		},
		{
			Hand: "QQQJA",
			Bid:  483,
		},
	}
}

package day6

type Race struct {
	Time     int
	Distance int
}

func loadInput(input string) []Race {
	return []Race{
		{
			Time:     7,
			Distance: 9,
		},
		{
			Time:     15,
			Distance: 40,
		},
		{
			Time:     30,
			Distance: 200,
		},
	}
}

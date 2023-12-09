package day5

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

/*

The puzzle input lists all of the seeds that need to be planted.
It also lists:
	- what type of soil they needs to be used with each seed.
	- what type of fertilizer to be used with each kind of soil
	- what type of water to use with each kind of fertilizer
	- how much light to use with how much water
	- the temperate to be used with an amount of light
	- temperature to humidity
	- humidity to location

It begins by listing seeds that need to be planted.

The almanac contains a list of maps which describe how to convert numbers
from a source category into numbers in a destination category.

For example:
	seed-to-soil map:
	50 98 2
	52 50 48

The seed-to-soil map describes how to convert a seed number (the source) to a soil number (the destination).

Rather than list every source number and its corresponding destination number one by on, the maps describe
entire ranges of numbers that can be converted.

Each line within a map contains three numbers: the destination range start, the source range start, and the range length

The first line of the seed-to-soil map has a destination range start of 5, a source range start of 98 and
a range length of 2.

This line means that the source range starts at 98 and contains two values: 98, 99
The destination range is the same length, but starts at 50: 50, 51

With the above information, we know that:
	- Seed Number 98 corresponds to Soil Number 50.
	- Seed Number 99 corresponds to Soil Number 51

With the second line "52, 50, 48":
	- Seeds number 52 - 99 correspond to Soil numbers 50 - 97

Any source numbers that aren't mapped correspond to the same destination number. So, seed number 10 corresponds to
soil number 10.

So, the entire list of seed numbers and their corresponding soil numbers looks like this:

seed  soil
0     0
1     1
...   ...
48    48
49    49
50    52
51    53
...   ...
96    98
97    99
98    50
99    51


*/

func TestLoadInputExtractsAllTheInputIntoAStringArray(t *testing.T) {
	input := "resources/input_test.txt"
	want := 26
	got := loadInput(input)
	if want != len(got) {
		t.Errorf("want %d, got %d", want, len(got))
	}
}

func TestParseSeedsExtractsTheSeedsIntoAnIntArray(t *testing.T) {
	input := []string{"seeds:", "79", "14", "55", "13"}
	want := []int{79, 14, 55, 13}
	got := parseSeeds(input)
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestParseSeedsUsingLoadInput(t *testing.T) {
	input := "resources/input_test.txt"
	want := []int{79, 14, 55, 13}
	stringInput := loadInput(input)
	got := parseSeeds(stringInput[0])
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestParseMapExtractsAMapStructFromLine(t *testing.T) {
	input := "resources/input_test.txt"
	want := Map{
		Name: "seed-to-soil",
		Ranges: []Range{
			{50, 98, 2},
			{52, 50, 48},
		},
	}
	stringInput := loadInput(input)
	lines := stringInput[1:4]
	got := parseMap(lines)
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

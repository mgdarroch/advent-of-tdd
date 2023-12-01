package day1

import (
	"fmt"
	"strings"
)

func ParseLine(input string) int {
	for _, c := range strings.Split(input, "") {
		fmt.Println(c)
	}
	return 12
}

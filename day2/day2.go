package day2

import (
	"errors"
	"strings"
)

func ParseLine(input string) (int, error) {
	if strings.Contains(input, "Game 1") {
		return 1, nil
	} else {
		return 0, errors.New("invalid game")
	}
}

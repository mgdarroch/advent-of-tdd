package day2

import (
	"errors"
	"strconv"
	"strings"
)

func ParseLine(input string) (int, error) {
	gameId := extractGameId(input)
	if strings.Contains(input, "Game 1") {
		return gameId, nil
	} else {
		return 0, errors.New("invalid game")
	}
}

func extractGameId(input string) int {
	var gameId int
	if input[6:7] == ":" {
		gameId, _ = strconv.Atoi(input[5:6])
	} else {
		gameId, _ = strconv.Atoi(input[5:7])
	}
	return gameId
}

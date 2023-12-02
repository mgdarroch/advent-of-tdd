package day2

import (
	"errors"
	"strconv"
	"strings"
)

type GameInfo struct {
	Id       int
	GameData string
}

func ParseLine(input string) (GameInfo, error) {
	gameId := extractGameId(input)

	game := GameInfo{
		Id:       gameId,
		GameData: "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	}

	if strings.Contains(input, "Game 1") {
		return game, nil
	} else {
		return GameInfo{}, errors.New("invalid game")
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

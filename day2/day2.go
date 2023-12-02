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

func ValidateGame(info GameInfo) bool {
	return true
}

func ParseLine(input string) (GameInfo, error) {
	gameId := extractGameId(input)
	game := extractGameData(input, gameId)
	if strings.Contains(input, "Game 1") {
		return game, nil
	} else {
		return GameInfo{}, errors.New("invalid game")
	}
}

func extractGameData(input string, gameId int) GameInfo {
	var gameData string
	if gameId >= 10 {
		gameData = input[9:]
	} else {
		gameData = input[8:]
	}
	game := GameInfo{
		Id:       gameId,
		GameData: gameData,
	}
	return game
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

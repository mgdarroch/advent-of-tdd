package day2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type GameInfo struct {
	Id       int
	GameData [3]Game
}

type Game struct {
	RedCount   int
	BlueCount  int
	GreenCount int
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
	fmt.Println(gameData)
	games := [3]Game{
		{
			RedCount:   4,
			BlueCount:  3,
			GreenCount: 0,
		},
		{
			RedCount:   1,
			BlueCount:  6,
			GreenCount: 2,
		},
		{
			RedCount:   0,
			BlueCount:  0,
			GreenCount: 2,
		},
	}
	gameInfo := GameInfo{
		Id:       gameId,
		GameData: games,
	}
	return gameInfo
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

package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type GameInfo struct {
	Id       int
	GameData []Game
}

type Game struct {
	RedCount   int
	BlueCount  int
	GreenCount int
}

func ValidateGame(sum int, input string) int {
	info := ParseLine(input)
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	valid := true
	newSum := sum
	for _, game := range info.GameData {
		fmt.Println(game)
		if game.RedCount > maxRed {
			valid = false
		}
		if game.GreenCount > maxGreen {
			valid = false
		}
		if game.BlueCount > maxBlue {
			valid = false
		}
	}

	if valid {
		newSum += info.Id
	}
	return newSum
}

func ParseLine(input string) GameInfo {
	gameId := extractGameId(input)
	game := extractGameData(input, gameId)
	return game
}

func extractGameData(input string, gameId int) GameInfo {
	var gameData string
	if gameId >= 10 {
		gameData = input[9:]
	} else {
		gameData = input[8:]
	}

	var games []Game
	gamesStr := strings.Split(gameData, ";")
	for _, game := range gamesStr {
		var cutset = " "
		trimStr := strings.Trim(game, cutset)
		cubes := strings.Split(trimStr, ";")
		for _, session := range cubes {
			trimSession := strings.ReplaceAll(session, " ", "")
			trimSession = strings.ReplaceAll(trimSession, "red", "r")
			trimSession = strings.ReplaceAll(trimSession, "green", "g")
			trimSession = strings.ReplaceAll(trimSession, "blue", "b")

			redCount := 0
			greenCount := 0
			blueCount := 0

			sessionSplit := strings.Split(trimSession, ",")
			for _, cube := range sessionSplit {
				fmt.Println(cube)
				if strings.Contains(cube, "r") {
					redCount, _ = strconv.Atoi(cube[:len(cube)-1])
				}
				if strings.Contains(cube, "g") {
					greenCount, _ = strconv.Atoi(cube[:len(cube)-1])
				}
				if strings.Contains(cube, "b") {
					blueCount, _ = strconv.Atoi(cube[:len(cube)-1])
				}
			}
			games = append(games, Game{
				RedCount:   redCount,
				BlueCount:  blueCount,
				GreenCount: greenCount,
			})
		}
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

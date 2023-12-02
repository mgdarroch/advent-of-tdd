package day2

import (
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

func ValidateGame(sum int, input string) (validSum int, cubePower int) {
	info := ParseLine(input)
	redLimit := 12
	greenLimit := 13
	blueLimit := 14

	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	valid := true
	newSum := sum
	for _, game := range info.GameData {
		if game.RedCount > redLimit {
			valid = false
		}
		if game.GreenCount > greenLimit {
			valid = false
		}
		if game.BlueCount > blueLimit {
			valid = false
		}
		if game.RedCount > maxRed {
			maxRed = game.RedCount
		}
		if game.GreenCount > maxGreen {
			maxGreen = game.GreenCount
		}
		if game.BlueCount > maxBlue {
			maxBlue = game.BlueCount
		}
	}

	if valid {
		newSum += info.Id
	}
	power := maxRed * maxBlue * maxGreen
	return newSum, power
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
	} else if gameId == 100 {
		gameData = input[7:]
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
	} else if input[7:8] == ":" {
		gameId, _ = strconv.Atoi(input[5:7])
	} else {
		gameId, _ = strconv.Atoi(input[5:8])
	}
	return gameId
}

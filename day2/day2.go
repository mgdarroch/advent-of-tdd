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

func Solve(input string, cubeSum int, powerSum int) (validSum int, cubePower int) {
	info := parseLine(input)
	redLimit, greenLimit, blueLimit := 12, 13, 14
	maxRed, maxGreen, maxBlue := 0, 0, 0

	valid := true
	newSum := cubeSum
	for _, game := range info.GameData {
		if game.RedCount > redLimit || game.GreenCount > greenLimit || game.BlueCount > blueLimit {
			valid = false
		}
		maxRed = max(maxRed, game.RedCount)
		maxGreen = max(maxGreen, game.GreenCount)
		maxBlue = max(maxBlue, game.BlueCount)
	}

	if valid {
		newSum += info.Id
	}
	newPower := powerSum + (maxRed * maxBlue * maxGreen)
	return newSum, newPower
}

func parseLine(input string) GameInfo {
	gameId := extractGameId(input)
	game := buildGameInfo(input, gameId)
	return game
}

func buildGameInfo(input string, gameId int) GameInfo {
	gameData := extractGameData(input, gameId)

	var games []Game
	gamesStr := strings.Split(gameData, ";")
	for _, game := range gamesStr {
		trimStr := strings.ReplaceAll(game, " ", "")
		cubes := strings.Split(trimStr, ";")
		for _, session := range cubes {
			trimSession := tidyInput(session)
			redCount, greenCount, blueCount := 0, 0, 0
			sessionSplit := strings.Split(trimSession, ",")
			for _, cube := range sessionSplit {
				colour := cube[len(cube)-1:]
				number, _ := strconv.Atoi(cube[:len(cube)-1])
				switch colour {
				case "b":
					blueCount = number
				case "g":
					greenCount = number
				case "r":
					redCount = number
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

func extractGameData(input string, gameId int) string {
	var gameData string
	if gameId >= 10 {
		gameData = input[9:]
	} else if gameId == 100 {
		gameData = input[7:]
	} else {
		gameData = input[8:]
	}
	return gameData
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

func tidyInput(session string) string {
	trimSession := strings.ReplaceAll(session, "red", "r")
	trimSession = strings.ReplaceAll(trimSession, "green", "g")
	trimSession = strings.ReplaceAll(trimSession, "blue", "b")
	return trimSession
}

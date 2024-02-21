package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"github.com/mitchellh/mapstructure"
)

type GameAttempt struct {
	Blue int `json:"blue"`
	Red int `json:"red"`
	Green int `json:"green"`
}

type GameResult struct {
	Game int `json:"game"`
	Attempts []GameAttempt `json:"attempts"`
}

func parseInput(line string) (*GameResult, error) {
	re := regexp.MustCompile(`Game (\d+):(.*)`)
	matches := re.FindAllStringSubmatch(line, -1)
	gameNumberString := matches[0][1]
	gameNumber, err := strconv.Atoi(gameNumberString)
	if err != nil {
		fmt.Printf("Error converting game number %s to integer: %v\n", gameNumberString, err)
		return nil, err
	}

	attemptsString := strings.TrimSpace(matches[0][2])
	attemptsArray := strings.Split(attemptsString, ";")

	gameResult := &GameResult{
		Game: gameNumber,
		Attempts: make([]GameAttempt, len(attemptsArray)),
	}

	for gameNumber, attempt := range attemptsArray {
		attempt := strings.TrimSpace(attempt)
		colorAndNumberStringsArray := strings.Split(attempt, ",")
		colorsMap := map[string]int{} 
		for _, colorAndNumberString := range colorAndNumberStringsArray {
			colorAndNumberString := strings.TrimSpace(colorAndNumberString)
			colorAndNumberArray := strings.Split(colorAndNumberString, " ")
			numberString, colorString := colorAndNumberArray[0], colorAndNumberArray[1]
			colorNumber, err := strconv.Atoi(numberString)
			if err != nil {
				fmt.Printf("Error converting game number %s to integer: %v\n", colorNumber, err)
				return nil, err
			}
			colorsMap[colorString] = colorNumber
		}
		var gameAttempt GameAttempt
		err := mapstructure.Decode(colorsMap, &gameAttempt)
		if err != nil {
			fmt.Println("Error: ", err)
			return nil, err
		}
		gameResult.Attempts[gameNumber] = gameAttempt
	}
	return gameResult, nil
}

func validateAttempt(attempt GameAttempt, attemptLimit GameAttempt) (bool) {
	return attempt.Blue <= attemptLimit.Blue && attempt.Red <= attemptLimit.Red && attempt.Green <= attemptLimit.Green
}

func validateGame(game GameResult, attemptLimit GameAttempt) (bool) {
	for _, attempt := range game.Attempts {
		isValid := validateAttempt(attempt, attemptLimit)
		if !isValid {
			return false
		}
	}
	return true;
}

func getTheMinimumNumberOfCubes(attempts []GameAttempt) (GameAttempt) {
	maxRed := 0
	maxBlue := 0
	maxGreen := 0
	for _, attempt := range attempts {
		if maxRed < attempt.Red  {
			maxRed = attempt.Red
		}
		if maxGreen < attempt.Green {
			maxGreen = attempt.Green
		}
		if maxBlue < attempt.Blue  {
			maxBlue = attempt.Blue
		}
	}

	return GameAttempt{
		Red: maxRed,
		Blue: maxBlue,
		Green: maxGreen,
	}
}


func partOne(file *os.File) {
	scanner := bufio.NewScanner(file)
	attemptLimit := GameAttempt{
		Blue: 14,
		Red: 12,
		Green: 13,
	}
	
	sumOfIds := 0
	for scanner.Scan() {
		line := scanner.Text()
		gameResult, err := parseInput(line)
		if err != nil {
			fmt.Println("Error parsing file:", err)
			return
		}
		isValid := validateGame(*gameResult, attemptLimit)
		if isValid {
			sumOfIds += gameResult.Game
		}
		fmt.Println(line)
		fmt.Println(gameResult.Game, isValid)
	}
	fmt.Println(sumOfIds)


	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumOfPowers := 0
	
	for scanner.Scan() {
		line := scanner.Text()
		gameResult, err := parseInput(line)
		if err != nil {
			fmt.Println("Error parsing file:", err)
			return
		}
		minAttempt := getTheMinimumNumberOfCubes(gameResult.Attempts)
		fmt.Println(line)
		fmt.Println(minAttempt)
		powerOfAttempt := minAttempt.Red * minAttempt.Blue * minAttempt.Green
		sumOfPowers += powerOfAttempt
	}

	fmt.Println(sumOfPowers)


	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}
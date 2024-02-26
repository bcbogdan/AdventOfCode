package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type CardGame struct {
	score          int
	winningNumbers map[string]bool
	gameNumbers    []string
}

func increaseGameScore(gameScore *int) {
	if *gameScore == 0 {
		*gameScore = 1
	} else {
		*gameScore = *gameScore * 2
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	linePattern := `^Card\s+(\d+):\s+(\d+(?:\s+\d+)*?) \|\s+(\d+(?:\s+\d+)*)$`
	lineRegexp := regexp.MustCompile(linePattern)
	numbersPatter := `\d+`
	numbersRegexp := regexp.MustCompile(numbersPatter)
	scanner := bufio.NewScanner(file)
	score := 0

	for scanner.Scan() {
		line := scanner.Text()
		matches := lineRegexp.FindStringSubmatch(line)
		cardNumber := matches[1]
		fmt.Printf("cardNumber: %v\n", cardNumber)
		winningNumbersString := matches[2]
		availableNumbersString := matches[3]
		winningNumbersMatches := numbersRegexp.FindAllString(winningNumbersString, -1)
		availableNumbersMatches := numbersRegexp.FindAllString(availableNumbersString, -1)

		gameScore := 0
		winningNumbers := map[string]bool{}
		for _, match := range winningNumbersMatches {
			winningNumbers[match] = true
		}
		for _, match := range availableNumbersMatches {
			_, hasMatch := winningNumbers[match]
			if hasMatch {
				// fmt.Printf("\"increase score\": %v\n", match)
				increaseGameScore(&gameScore)
			}
		}

		score += gameScore
		fmt.Printf("gameScore: %v\n", gameScore)
	}
	fmt.Printf("score: %v\n", score)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}

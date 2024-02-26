package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type ScratchCard struct {
	number  string
	matches int
}

func recursiveCount(index int, scratchCard *ScratchCard, allScratchCards *[]ScratchCard) int {
	count := 1
	if scratchCard.matches == 0 {
		return count
	}

	lastMatchIndex := index + scratchCard.matches
	nextCardIndex := index + 1

	// fmt.Printf("lastMatchIndex: %v\n", lastMatchIndex)
	// fmt.Printf("nextCardIndex: %v\n", nextCardIndex)
	for nextCardIndex <= lastMatchIndex {
		if nextCardIndex >= len(*allScratchCards) {
			break
		}
		count += recursiveCount(nextCardIndex, &(*allScratchCards)[nextCardIndex], allScratchCards)
		nextCardIndex++
	}

	return count
}
func countScratchCards(allScratchCards *[]ScratchCard) int {
	count := 0

	for index, scratchCard := range *allScratchCards {
		fmt.Printf("scratchCard: %v\n", scratchCard)
		count += recursiveCount(index, &scratchCard, allScratchCards)
	}
	return count
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

	scratchCards := []ScratchCard{}

	for scanner.Scan() {
		line := scanner.Text()
		matches := lineRegexp.FindStringSubmatch(line)
		cardNumber := matches[1]
		fmt.Printf("cardNumber: %v\n", cardNumber)
		scratchCard := ScratchCard{
			number: cardNumber,
		}
		winningNumbersString := matches[2]
		availableNumbersString := matches[3]
		winningNumbersMatches := numbersRegexp.FindAllString(winningNumbersString, -1)
		availableNumbersMatches := numbersRegexp.FindAllString(availableNumbersString, -1)
		winningNumbers := map[string]bool{}
		for _, match := range winningNumbersMatches {
			winningNumbers[match] = true
		}
		for _, match := range availableNumbersMatches {
			_, hasMatch := winningNumbers[match]
			if hasMatch {
				scratchCard.matches += 1
			}
		}
		// fmt.Printf("scratchCard: %v\n", scratchCard)
		scratchCards = append(scratchCards, scratchCard)

	}
	count := countScratchCards(&scratchCards)
	fmt.Printf("score: %v\n", count)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}

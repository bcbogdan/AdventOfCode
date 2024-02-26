package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type SymbolsMap struct {
	previousRow map[int]rune
	currentRow  map[int]rune
}

type PartNumber struct {
	number     int
	startIndex int
	endIndex   int
	rowIndex   int
}

func (partNumber PartNumber) getIndexes() []int {
	size := partNumber.endIndex - partNumber.startIndex + 1

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = partNumber.startIndex + i
	}
	return arr
}

func (partNumber PartNumber) isValid(symbolsMap *SymbolsMap) bool {
	startIndex := partNumber.startIndex - 1
	endIndex := partNumber.endIndex + 1

	hasAdjacentSymbol := false

	for startIndex <= endIndex {
		_, prevRowOk := symbolsMap.previousRow[startIndex]
		if prevRowOk {
			hasAdjacentSymbol = true
			break
		}
		_, currentRowOk := symbolsMap.currentRow[startIndex]
		if currentRowOk {
			hasAdjacentSymbol = true
			break
		}
		startIndex += 1
	}
	return hasAdjacentSymbol
}

type Gear struct {
	x           int
	y           int
	partNumbers []PartNumber
}

func (gear *Gear) getRatio() int {
	return gear.partNumbers[0].number * gear.partNumbers[1].number
}

func (gear *Gear) hasPartNumber(partNumber *PartNumber) bool {
	hasPartNumber := false
	for _, pN := range gear.partNumbers {
		if pN.endIndex == partNumber.endIndex && pN.startIndex == partNumber.startIndex && pN.rowIndex == partNumber.rowIndex {
			hasPartNumber = true
			break
		}
	}
	return hasPartNumber
}

func (gear *Gear) process(partNumbersRowMap map[int]PartNumber) {
	indexesToCheck := [3]int{gear.x - 1, gear.x, gear.x + 1}

	if len(gear.partNumbers) == 2 {
		return
	}
	for _, index := range indexesToCheck {
		partNumber, hasAdjacentRowNumber := partNumbersRowMap[index]

		if hasAdjacentRowNumber && !gear.hasPartNumber(&partNumber) {
			gear.partNumbers = append(gear.partNumbers, partNumber)
		}
		if len(gear.partNumbers) == 2 {
			break
		}
	}
}

func (gear *Gear) isValid() bool {
	return len(gear.partNumbers) == 2
}

type RuneEvaluator struct {
	currentRowIndex int
	stringNumber    string
	symbols         SymbolsMap
	// unvalidatedPartNumbers []PartNumber
	// validatedPartNumbers   []PartNumber
	previousRowPartNumbers []PartNumber
	currentRowPartNumbers  []PartNumber
	validatedGears         []Gear
	unvalidatedGears       []Gear
}

func (evaluator *RuneEvaluator) evaluate(character rune, characterIndex int) {
	isDigit := unicode.IsDigit(character)
	if isDigit {
		evaluator.stringNumber = evaluator.stringNumber + string(character)
		return
	}

	if character == '*' {
		gear := &Gear{
			x:           characterIndex,
			y:           evaluator.currentRowIndex,
			partNumbers: []PartNumber{},
		}
		evaluator.unvalidatedGears = append(evaluator.unvalidatedGears, *gear)
	}

	isSymbol := character != '.'

	if isSymbol {
		evaluator.symbols.currentRow[characterIndex] = character
	}

	if evaluator.stringNumber == "" {
		return
	}

	evaluator.convertStringNumber(characterIndex)
}

func (evaluator *RuneEvaluator) init() {
	evaluator.symbols.currentRow = make(map[int]rune)
	evaluator.symbols.previousRow = make(map[int]rune)
}

func (evaluator *RuneEvaluator) nextRow(lineLength int) {
	if evaluator.stringNumber != "" {
		evaluator.convertStringNumber(lineLength)
	}

	// fmt.Printf("evaluator.validatedGears: %v\n", evaluator.validatedGears)
	// fmt.Printf("evaluator.unvalidatedGears: %v\n", evaluator.unvalidatedGears)
	// fmt.Printf("evaluator.currentRowIndex: %v\n", evaluator.currentRowIndex)
	currentRowPartNumbersMap := map[int]PartNumber{}
	prevRowPartNumbersMap := map[int]PartNumber{}

	for _, partNumber := range evaluator.previousRowPartNumbers {
		indexes := partNumber.getIndexes()
		for _, index := range indexes {
			prevRowPartNumbersMap[index] = partNumber
		}
	}

	for _, partNumber := range evaluator.currentRowPartNumbers {
		indexes := partNumber.getIndexes()
		for _, index := range indexes {
			currentRowPartNumbersMap[index] = partNumber
		}
	}

	unvalidatedGears := []Gear{}
	for _, gear := range evaluator.unvalidatedGears {
		if gear.y > evaluator.currentRowIndex-2 {
			unvalidatedGears = append(unvalidatedGears, gear)
		}
	}
	evaluator.unvalidatedGears = unvalidatedGears
	unvalidatedGears = []Gear{}
	for _, gear := range evaluator.unvalidatedGears {
		gear.process(prevRowPartNumbersMap)
		gear.process(currentRowPartNumbersMap)
		if gear.isValid() {
			evaluator.validatedGears = append(evaluator.validatedGears, gear)
		} else {
			unvalidatedGears = append(unvalidatedGears, gear)
		}
	}
	evaluator.unvalidatedGears = unvalidatedGears

	// filteredPartNumbers := []PartNumber{}
	// for _, partNumber := range evaluator.unvalidatedPartNumbers {
	// 	if partNumber.rowIndex > evaluator.currentRowIndex-2 {
	// 		filteredPartNumbers = append(filteredPartNumbers, partNumber)
	// 	}
	// }
	// evaluator.unvalidatedPartNumbers = filteredPartNumbers
	// filteredPartNumbers = []PartNumber{}

	// for _, partNumber := range evaluator.unvalidatedPartNumbers {
	// 	if partNumber.isValid(&evaluator.symbols) {
	// 		evaluator.validatedPartNumbers = append(evaluator.validatedPartNumbers, partNumber)
	// 	} else {
	// 		filteredPartNumbers = append(filteredPartNumbers, partNumber)
	// 	}
	// }
	// evaluator.unvalidatedPartNumbers = filteredPartNumbers

	evaluator.currentRowIndex += 1
	evaluator.symbols.previousRow = evaluator.symbols.currentRow
	evaluator.symbols.currentRow = make(map[int]rune)

	evaluator.previousRowPartNumbers = evaluator.currentRowPartNumbers
	evaluator.currentRowPartNumbers = []PartNumber{}
}

func (evaluator *RuneEvaluator) convertStringNumber(index int) {
	convertedNumber, err := strconv.Atoi(evaluator.stringNumber)
	if err != nil {
		fmt.Printf("convert err: %v\n", err)
	} else {

		partNumber := &PartNumber{
			number:     convertedNumber,
			startIndex: index - len(evaluator.stringNumber),
			endIndex:   index - 1,
			rowIndex:   evaluator.currentRowIndex,
		}
		// evaluator.unvalidatedPartNumbers = append(evaluator.unvalidatedPartNumbers, *partNumber)
		evaluator.currentRowPartNumbers = append(evaluator.currentRowPartNumbers, *partNumber)
	}

	evaluator.stringNumber = ""
}

func (evaluator *RuneEvaluator) sumOfPartNumbers() int {
	sum := 0

	// for _, partNumber := range evaluator.validatedPartNumbers {
	// 	fmt.Printf("partNumber: %v\n", partNumber)
	// 	sum += partNumber.number
	// }

	return sum
}

func (evaluator *RuneEvaluator) sumOfGearRatios() int {
	sum := 0

	for _, gear := range evaluator.validatedGears {
		fmt.Printf("gearRatio: %v\n", gear)
		sum += gear.getRatio()
	}

	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowIndex := -1
	runeEvaluator := &RuneEvaluator{}
	runeEvaluator.init()
	for scanner.Scan() {
		rowIndex += 1
		line := scanner.Text()
		// fmt.Printf("runeEvaluator.currentRowIndex: %v\n", runeEvaluator.currentRowIndex)
		// fmt.Printf("runeEvaluator.unvalidatedPartNumbers: %v\n", runeEvaluator.unvalidatedPartNumbers)
		// fmt.Printf("runeEvaluator.validatedPartNumbers: %v\n", runeEvaluator.validatedPartNumbers)
		// fmt.Printf("runeEvaluator.symbols.currentRow: %v\n", runeEvaluator.symbols.currentRow)
		// fmt.Printf("runeEvaluator.symbols.previousRow: %v\n", runeEvaluator.symbols.previousRow)
		for index, character := range line {
			runeEvaluator.evaluate(character, index)
		}
		runeEvaluator.nextRow(len(line))
	}

	// fmt.Printf("runeEvaluator.sumOfPartNumbers(): %v\n", runeEvaluator.sumOfPartNumbers())
	fmt.Printf("runeEvaluator.sumOfGearRatios(): %v\n", runeEvaluator.sumOfGearRatios())
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}

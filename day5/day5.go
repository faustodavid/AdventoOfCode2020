package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	inputString := string(input)
	puzzle1Answer := 0
	var seats [128 * 8]bool
	var min, max int

	hasNewLine := true
	for hasNewLine {
		indexOfBreakLine := strings.Index(inputString, "\n")
		if indexOfBreakLine == -1 {
			hasNewLine = false
			indexOfBreakLine = len(inputString)
		} else {
			inputString = inputString[indexOfBreakLine+1:]
		}

		seatId := GetSeatId(inputString[:indexOfBreakLine])

		seats[seatId] = true
		if seatId < min || min == 0 {
			min = seatId
		}
		if seatId > max {
			max = seatId
		}

		if seatId > puzzle1Answer {
			puzzle1Answer = seatId
		}
	}

	puzzle2Answer := 0
	for id, seatTaken := range seats[min+1 : max] {
		if !seatTaken {
			puzzle2Answer = id + min + 1
		}
	}

	fmt.Printf("Puzzle1: %d\n", puzzle1Answer)
	fmt.Printf("Puzzle2: %d\n", puzzle2Answer)
}

func GetSeatId(boardingId string) int {
	return (GetRow(boardingId) * 8) + GetColumn(boardingId)
}

func GetRow(boardingId string) int {
	rowF := 0
	rowB := 127

	for i := 0; i < 6; i++ {
		if boardingId[i] == 'F' {
			rowB = (rowF + rowB) / 2
		} else {
			rowF = ((rowB - rowF) / 2) + rowF
		}
	}

	if boardingId[6] == 'F' {
		return rowF + 1
	} else {
		return rowB
	}
}

func GetColumn(boardingId string) int {
	columnUpper := 7
	columnLower := 0

	for i := 7; i < 9; i++ {
		if boardingId[i] == 'R' {
			columnLower = ((columnUpper + columnLower) / 2) + 1
		} else {
			columnUpper = ((columnLower - columnUpper) / 2) + columnUpper - 1
		}
	}

	if boardingId[9] == 'R' {
		return columnUpper
	} else {
		return columnLower
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("/Users/fsuarezrosario/go/src/AdventoOfCode2020/day3/input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	inputString := string(input)
	grid := StringToByteGrid(inputString)

	print("Puzzle1: ")
	println(CountTreesOnGridUsingPattern(grid, 3, 1))

	puzzle2Answer := CountTreesOnGridUsingPattern(grid, 1, 1) *
		CountTreesOnGridUsingPattern(grid, 3, 1) *
		CountTreesOnGridUsingPattern(grid, 5, 1) *
		CountTreesOnGridUsingPattern(grid, 7, 1) *
		CountTreesOnGridUsingPattern(grid, 1, 2)

	print("Puzzle2: ")
	println(puzzle2Answer)
}

func StringToByteGrid(input string) [][]byte {
	linesCount := strings.Count(input, "\n")
	lineLength := strings.Index(input, "\n")

	var grid [][]byte
	for i := 0; i <= linesCount; i++ {
		line := make([]byte, lineLength)
		for n := 0; n < lineLength; n++ {
			line[n] = input[n+(lineLength*i)+i]
		}
		grid = append(grid, line)
	}

	return grid
}

func CountTreesOnGridUsingPattern(tree [][]byte, right int, down int) int {
	rowLength := len(tree[0])
	rowReplicas := 0
	currentRightPosition := 0
	currentDownPosition := 0
	treeCount := 0
	treeByte := byte('#')

	for true {
		currentRightPosition += right
		currentDownPosition += down

		if currentDownPosition >= len(tree) {
			break
		}

		for currentRightPosition-(rowLength*rowReplicas) >= rowLength {
			rowReplicas++
		}

		if tree[currentDownPosition][currentRightPosition-(rowLength*rowReplicas)] == treeByte {
			treeCount++
		}

	}
	return treeCount
}

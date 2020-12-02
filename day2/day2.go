package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	inputString := string(input)
	validPasswordsPuzzle1Count := 0
	validPasswordsPuzzle2Count := 0

	for true {
		// Look for min
		index := strings.Index(inputString, "-")
		min, _ := strconv.Atoi(inputString[:index])
		inputString = inputString[index+1:]

		// Look for max
		index = strings.Index(inputString, " ")
		max, _ := strconv.Atoi(inputString[:index])
		inputString = inputString[index+1:]

		// Look for keyword
		index = strings.Index(inputString, ":")
		keyword := inputString[:index]
		inputString = inputString[index+2:]

		// Look for password
		index = strings.Index(inputString, "\n")
		if index > -1 {
			password := inputString[:index]

			if IsValidPasswordPuzzle1(password, keyword, min, max) {
				validPasswordsPuzzle1Count++
			}

			if IsValidPasswordPuzzle2(password, keyword[0], min, max) {
				validPasswordsPuzzle2Count++
			}

			inputString = inputString[index+1:]

		} else {
			password := inputString

			if IsValidPasswordPuzzle1(password, keyword, min, max) {
				validPasswordsPuzzle1Count++
			}

			if IsValidPasswordPuzzle2(password, keyword[0], min, max) {
				validPasswordsPuzzle2Count++
			}

			break
		}
	}

	fmt.Printf("Puzzle1: %d\n", validPasswordsPuzzle1Count)
	fmt.Printf("Puzzle2: %d\n", validPasswordsPuzzle2Count)
}

func IsValidPasswordPuzzle1(password string, keyword string, min int, max int) bool {
	keyCount := 0

	for true {
		index := strings.Index(password, keyword)
		if index > -1 {
			keyCount++

			if keyCount > max {
				return false
			}

			password = password[index+1:]
		} else {
			break
		}
	}

	return keyCount <= max && keyCount >= min
}

func IsValidPasswordPuzzle2(password string, keyword byte, firstPosition int, secondPosition int) bool {
	return (password[firstPosition-1] == keyword && password[secondPosition-1] != keyword) ||
		(password[firstPosition-1] != keyword && password[secondPosition-1] == keyword)
}

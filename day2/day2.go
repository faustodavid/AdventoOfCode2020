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
	count := 0

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
		inputString = inputString[index+1:]

		// Look for password
		index = strings.Index(inputString, "\n")
		if index > -1 {
			password := inputString[:index]

			if IsPasswordValid(CorruptedPassword{password, keyword, min, max}) {
				count++
			}

			inputString = inputString[index+1:]

		} else {
			password := inputString

			if IsPasswordValid(CorruptedPassword{password, keyword, min, max}) {
				count++
			}

			break
		}
	}

	fmt.Print("Puzzle1:")
	fmt.Println(count)
}

type CorruptedPassword struct {
	password string
	keyword  string
	min      int
	max      int
}

func IsPasswordValid(corruptedPassword CorruptedPassword) bool {
	keyCount := 0

	for true {
		index := strings.Index(corruptedPassword.password, corruptedPassword.keyword)
		if index > -1 {
			keyCount++

			if keyCount > corruptedPassword.max {
				return false
			}

			corruptedPassword.password = corruptedPassword.password[index+1:]
		} else {
			break
		}
	}

	return keyCount <= corruptedPassword.max && keyCount >= corruptedPassword.min
}

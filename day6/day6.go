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

	answers := string(input)

	puzzle1 := CountUniqueLetters(GroupAnswersByGroups(answers))

	fmt.Printf("Puzzle1: %d\n", puzzle1)
}

func GroupAnswersByGroups(answers string) []string {
	groups := make([]string, strings.Count(answers, "\n\n")+1)

	for i := 0; i < len(groups)-1; i++ {
		breakLineIndex := strings.Index(answers, "\n\n")
		group := answers[:breakLineIndex]
		group = strings.Replace(group, "\n", "", -1)
		groups[i] = group

		answers = answers[breakLineIndex+2:]
	}

	groups[len(groups)-1] = answers

	return groups
}

func CountUniqueLetters(groupedAnswers []string) int {
	count := 0
	for _, answer := range groupedAnswers {
		uniqueLetters := make(map[int32]bool)
		for _, letter := range answer {
			if letter == 10 {
				continue
			}
			uniqueLetters[letter] = true
		}

		count += len(uniqueLetters)
	}

	return count
}

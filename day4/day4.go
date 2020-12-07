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

	puzzle1, puzzle2 := countValidPassports(inputString)

	println("Puzzle1: %d", puzzle1)
	println("Puzzle2: %d", puzzle2)
}

func countValidPassports(input string) (int, int) {
	validPassportsPuzzle1Count := 0
	validPassportsPuzzle2Count := 0

	for true {
		newLineIndex := strings.Index(input, "\n\n")
		if newLineIndex == -1 {
			passport := input
			if isPassportValidPuzzle1(passport) {
				validPassportsPuzzle1Count++
			}
			if isPassportValidPuzzle2(passport) {
				validPassportsPuzzle2Count++
			}
			break
		}

		passport := input[:newLineIndex]
		if isPassportValidPuzzle1(passport) {
			validPassportsPuzzle1Count++
		}
		if isPassportValidPuzzle2(passport) {
			validPassportsPuzzle2Count++
		}

		input = input[newLineIndex+2:]
	}

	return validPassportsPuzzle1Count, validPassportsPuzzle2Count
}

func isPassportValidPuzzle1(passport string) bool {
	propertiesCount := strings.Count(passport, " ") + strings.Count(passport, "\n") + 1

	if propertiesCount == 8 {
		return true
	}

	return propertiesCount == 7 && !strings.Contains(passport, "cid")
}

func isPassportValidPuzzle2(passport string) bool {
	containsCid := strings.Contains(passport, "cid")
	propertiesCount := 0
	for true {
		indexOfBlankSpace := strings.Index(passport, " ")
		indexOfBreakLine := strings.Index(passport, "\n")
		indexOfPropertyEnd := -1

		if (indexOfBlankSpace < indexOfBreakLine && indexOfBlankSpace > -1) || indexOfBreakLine == -1 {
			indexOfPropertyEnd = indexOfBlankSpace
		} else {
			indexOfPropertyEnd = indexOfBreakLine
		}

		if indexOfPropertyEnd == -1 {
			property := passport
			if !isPropertyValid(property) {
				return false
			}

			propertiesCount++
			break
		}

		property := passport[:indexOfPropertyEnd]
		passport = passport[indexOfPropertyEnd+1:]

		if !isPropertyValid(property) {
			return false
		}

		propertiesCount++
	}

	return propertiesCount == 8 || (propertiesCount == 7 && !containsCid)
}

func isPropertyValid(property string) bool {
	// For cleaner way we can use split, but to avoid allocations lets micro optimize this :)
	separatorIndex := strings.Index(property, ":")
	propertyName := property[:separatorIndex]
	propertyValue := property[separatorIndex+1:]

	switch propertyName {
	case "byr":
		byr, _ := strconv.Atoi(propertyValue)
		if byr >= 1920 && byr <= 2002 {
			return true
		}

	case "iyr":
		iyr, _ := strconv.Atoi(propertyValue)
		if iyr >= 2010 && iyr <= 2020 {
			return true
		}

	case "eyr":
		eyr, _ := strconv.Atoi(propertyValue)
		if eyr >= 2020 && eyr <= 2030 {
			return true
		}

	case "hgt":
		hgt, _ := strconv.Atoi(propertyValue[:len(propertyValue)-2])
		if propertyValue[len(propertyValue)-2:] == "cm" {
			if hgt >= 150 && hgt <= 193 {
				return true
			}
		} else {
			if hgt >= 59 && hgt <= 76 {
				return true
			}
		}

	case "hcl":
		return isValidHcl(propertyValue)
	case "ecl":
		return propertyValue == "amb" ||
			propertyValue == "blu" ||
			propertyValue == "brn" ||
			propertyValue == "gry" ||
			propertyValue == "grn" ||
			propertyValue == "hzl" ||
			propertyValue == "oth"
	case "pid":
		return isValidPid(propertyValue)
	case "cid":
		return true
	default:
		return false
	}

	return false
}

func isValidPid(pid string) bool {
	if len(pid) != 9 {
		return false
	}

	for i := 0; i < 9; i++ {
		current := pid[i]
		if current >= '0' && current <= '9' {
			continue
		}

		return false
	}

	return true
}

func isValidHcl(hcl string) bool {
	if hcl[0] != "#"[0] {
		return false
	}

	if len(hcl) != 7 {
		return false
	}

	alphaNumeric := hcl[1:]
	for i := 0; i < len(alphaNumeric); i++ {
		current := alphaNumeric[i]
		if current >= '0' && current <= '9' {
			continue
		}
		if current >= 'a' && current <= 'f' {
			continue
		}

		return false
	}

	return true
}

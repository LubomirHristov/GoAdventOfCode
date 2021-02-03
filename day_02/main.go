package main

import (
	"fmt"
	"strconv"
	"strings"

	"common"
)

func isValidPolicy1(min int, max int, letter rune, password string) bool {
	occurences := 0

	for _, currentLetter := range password {
		if letter == currentLetter {
			occurences++
		}
	}

	return min <= occurences && occurences <= max
}

func isValidPolicy2(min int, max int, letter rune, password string) bool {
	x := rune(password[min-1]) == letter
	y := rune(password[max-1]) == letter

	return x != y
}

func findValidPasswords(data []string) int {
	validPasswords := 0

	for _, line := range data {
		parts := strings.Fields(line)

		minAndMax := strings.Split(parts[0], "-")

		min, _ := strconv.Atoi(minAndMax[0])
		max, _ := strconv.Atoi(minAndMax[1])
		letter := rune(strings.Trim(parts[1], ":")[0])
		password := parts[2]

		if isValidPolicy2(min, max, letter, password) {
			validPasswords++
		}

	}

	return validPasswords
}

func main() {
	data := common.ReadFileString("input.txt")
	fmt.Println(findValidPasswords(data))
}

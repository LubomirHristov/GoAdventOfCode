package main

import (
	"fmt"
	"strings"

	"common"
)

func treesOnTheWay(data []string, right int, down int) int {
	count := 0
	pos := 0

	for index, line := range data {
		if !(index%down == 0) {
			continue
		}

		if pos >= len(line) {
			line = strings.Repeat(line, 100)

		}

		if line[pos] == '#' {
			count++
		}

		pos += right

	}

	return count
}

func main() {
	data := common.ReadFileString("input.txt")
	slope1 := treesOnTheWay(data, 1, 1)
	slope2 := treesOnTheWay(data, 3, 1)
	slope3 := treesOnTheWay(data, 5, 1)
	slope4 := treesOnTheWay(data, 7, 1)
	slope5 := treesOnTheWay(data, 1, 2)

	fmt.Println(slope1 * slope2 * slope3 * slope4 * slope5)
}

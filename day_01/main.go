package main

import (
	"fmt"

	"common"
)

func findProduct(data []int, upto int) int {
	remaining := make(map[int]bool)

	for _, num := range data {
		if _, ok := remaining[num]; ok {
			return num * (upto - num)
		}

		remaining[upto-num] = true
	}

	return -1
}

func main() {
	data := common.ReadFileInt("input.txt")
	// fmt.Println(findProduct(data, 2020))
	for _, num := range data {
		if res := findProduct(data, 2020-num); res != -1 {
			fmt.Println(num * res)
		}
	}
}

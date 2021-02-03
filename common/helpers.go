package common

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFileInt(pathToFile string) []int {
	inputArray := make([]int, 0)

	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in, _ := strconv.Atoi(scanner.Text())
		inputArray = append(inputArray, in)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return inputArray
}

func ReadFileString(pathToFile string) []string {
	inputArray := make([]string, 0)

	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputArray = append(inputArray, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return inputArray
}

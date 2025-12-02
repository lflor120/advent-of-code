package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
)

// thought process: this one is gonna be ugly, 
// I'm thinking loop through each range and use string parsing to see if numbers are repeated

func main() {
	fmt.Println("Day 2 2025!")
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	ranges := processFile(string(file))

	fmt.Println(ranges)
}

func processFile(fileContents string) [][]int {
	ranges := make([][]int, 0)
	fileRanges := strings.Split(fileContents, ",")

	for _, fileRange := range fileRanges {
		fileNums := strings.Split(fileRange, "-")
		numOne, err := strconv.Atoi(fileNums[0])
		if err != nil {
			log.Fatalf("failed to convert %v to int!. err: %v", numOne, err)
		}

		numTwo, err := strconv.Atoi(fileNums[1])
		if err != nil {
			log.Fatalf("failed to convert %v to int!. err: %v", numTwo, err)
		}
		ranges = append(ranges, []int{numOne, numTwo})
	}
	return ranges
}

func isRepeatedNumber(num int) bool {
	numStr := strconv.Itoa(num)
	halfway := len(numStr) / 2
	return numStr[:halfway] != numStr[halfway+1:]
}
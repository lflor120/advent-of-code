package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Day 2!")
	fileContents, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	reports := getReports(string(fileContents))


}

func getReports(fileContents string) [][]int {
	reports := make([][]int, 0)

	return reports
}

func isIncreasing(slice []int) bool {
	return true
}

func isDecreasing(slice []int) bool {
	return true
}

func validDifference(slice []int) bool {
	// returns whether adjacent elements differ by at least one and at most three
	return true
}

func isSafe(slice []int) bool {
	monotonic := isIncreasing(slice) || isDecreasing(slice)
	diff := validDifference(slice)
	return monotonic && diff
}
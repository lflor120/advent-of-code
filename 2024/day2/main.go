package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 2!")
	fileContents, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	reports := getReports(string(fileContents))
	safeReports := getSafeReports(reports)
	fmt.Printf("Solution one: %d \n", safeReports)
}

func getReports(fileContents string) [][]int {
	// converts input.txt to a 2D array of ints
	reports := make([][]int, 0)

	lines := strings.Split(fileContents, "\n")
	for _, line := range lines {
		levelsString := strings.Fields(line)
		levels := make([]int, 0)
		for _, level := range levelsString {
			num, err := strconv.Atoi(level)
			if err != nil {
				log.Fatalf("Failed to convert %d to an int. err: %v", num, err)
			}
			levels = append(levels, num)
		}

		reports = append(reports, levels)
	}
	return reports
}

func isIncreasing(slice []int) bool {
	for i := range len(slice) - 1 {
		if slice[i] > slice[i + 1] {
			return false
		}
	}
	return true
}

func isDecreasing(slice []int) bool {
	for i := range len(slice) - 1 {
		if slice[i] < slice[i + 1] {
			return false
		}
	}
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

func getSafeReports(reports [][]int) int {
	// gets the number of safe reports
	safeReports := 0
	for _, report := range reports {
		if isSafe(report) {
			safeReports++
		}
	}
	return safeReports
}
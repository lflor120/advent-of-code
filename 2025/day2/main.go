package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	solutionOne := getAllInvalidIds(ranges)
	fmt.Printf("Solution one: %d", solutionOne)
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

func isRepeatedTwice(num int) bool {
	numStr := strconv.Itoa(num)
	halfway := len(numStr) / 2
	return numStr[:halfway] == numStr[halfway:] && (len(numStr)%2 == 0)
}

func part2BruteForce(num int) bool {
	numStr := strconv.Itoa(num)

	for j := 0; j < len(numStr); j++ {
		window := numStr[:j]
		transformed := strings.ReplaceAll(numStr, window, "")

		if transformed == "" && len(window) != len(numStr) && strings.Count(numStr, window) > 1 {
			return true
		}
	}
	fmt.Println("----")
	return false
}

func getInvalidIdsInRange(start, stop int) []int {
	invalidIds := make([]int, 0)
	for i := start; i <= stop; i++ {
		if isInvalid(i) {
			invalidIds = append(invalidIds, i)
		}
	}
	return invalidIds
}

func isInvalid(num int) bool {
	return part2BruteForce(num)
	// numStr := strconv.Itoa(num)
	// cleanStr := removeAdjacent(numStr)
	// cleanNum, err := strconv.Atoi(cleanStr)
	// if err != nil {
	// 	log.Fatal("Sad")
	// }
	// return isRepeatedTwice(num) || isRepeatedMultiple(num) ||  isRepeatedMultiple(cleanNum)
}

func removeAdjacent(str string) string {
	clean := ""
	for i := 0; i < len(str); i++ {
		if i > 0 && str[i-1] == str[i] {
			continue
		}
		clean += string(str[i])
	}
	return clean
}

func isRepeatedMultiple(num int) bool {
	numStr := strconv.Itoa(num)
	window := ""
	seen := map[rune]bool{} // this will act as a set

	for _, c := range numStr {
		// expand window
		if _, ok := seen[c]; !ok {
			window += string(c)
			seen[c] = true
			continue
		}

		// check if repeated length checks out
		transformed := strings.ReplaceAll(numStr, window, "")
		return transformed == "" && len(window) != len(numStr)
	}
	return false
}

func getAllInvalidIds(ranges [][]int) int {
	total := 0

	for _, idRange := range ranges {
		start, stop := idRange[0], idRange[1]
		fmt.Printf("Checking range: %d - %d\n", start, stop)
		invalidIds := getInvalidIdsInRange(start, stop)
		fmt.Printf("invalid ids: %v\n", invalidIds)
		fmt.Println("------")
		total += sum(invalidIds)
	}
	return total
}

func sum(array []int) int {
	total := 0
	for _, num := range array {
		total += num
	}
	return total
}

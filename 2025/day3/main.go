package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 3, 2025!")
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to open file! %v", err)
	}
	banks := getBanks(string(file))

	solutionOne := getTotalJoltage(banks)
	fmt.Printf("Solution one: %d\n", solutionOne)

	solutionTwo := getTotalNJoltage(banks, 12)
	fmt.Printf("Solution two: %d", solutionTwo)
}

func getBanks(fileContents string) [][]int {
	banks := strings.Split(fileContents, "\n")
	bankInts := make([][]int, 0)

	for _, bank := range banks {
		currentBank := make([]int, 0)
		for _, num := range bank {
			numInt, err := strconv.Atoi(string(num))
			if err != nil {
				log.Fatalf("failed to convert: %v to int!", num)
			}

			currentBank = append(currentBank, numInt)
		}
		bankInts = append(bankInts, currentBank)
	}
	return bankInts
}

func getLargestJoltage(bank []int) int {
	largestJolt := 0
	for i, left := range bank {
		for _, right := range bank[i+1:] {
			currentJolt := (left * 10) + right
			largestJolt = max(largestJolt, currentJolt)
		}
	}
	return largestJolt
}

func getLargestNJoltage(bank []int, n int) int64 {
	jolts := make([]int, 0)
	// prevIndex := 0
	want := n
	total := len(bank)
	start := 0

	for want > 0 {
		end := total - (want - 1)
		if want == 1 {
			end = total
		}

		window := bank[start:end]
		currMax, maxIndex := getLeftMostMax(window) 
		jolts = append(jolts, currMax)

		start = maxIndex + start + 1      
		want--                        
	}

	largestJolt := convertSliceToInt(jolts)
	return largestJolt
}

func convertSliceToInt(slice []int) int64 {
	sliceStr := ""
	for _, num := range slice {
		sliceStr += strconv.Itoa(num)
	}

	sliceInt64, err := strconv.ParseInt(sliceStr, 10, 64)
	if err != nil {
		log.Fatalf("failed to convert %v to int! err: %v", sliceStr, err)
	}
	return sliceInt64
}

// returns the num and the index
func getLeftMostMax(bank []int) (int, int) {
	leftMax := 0
	maxIndex := 0
	for i, num := range bank {
		if num > leftMax {
			leftMax = num
			maxIndex = i
		}
	}
	return leftMax, maxIndex
}

func getTotalJoltage(banks [][]int) int {
	var joltage int
	for _, bank := range banks {
		joltage += getLargestJoltage(bank)
	}
	return joltage
}

func getTotalNJoltage(banks [][]int, n int) int64 {
	var joltage int64
	for _, bank := range banks {
		joltage += getLargestNJoltage(bank, n)
	}
	return joltage
}

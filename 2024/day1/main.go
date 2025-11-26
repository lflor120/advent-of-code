package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 1!")
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("failed to open file!")
	}

	lines := strings.Split(string(file), "\n")
	firstSlice, secondSlice := convertToSortedIntSlice(lines)

	solutionOne := getDistances(firstSlice, secondSlice)
	fmt.Printf("Solution one: %d \n", solutionOne)

	solutionTwo := getSimilarityScore(firstSlice, secondSlice)
	fmt.Printf("Solution two: %d", solutionTwo)
}

func convertToSortedIntSlice(lines []string) ([]int, []int) {
	firstSlice := make([]int, 0)
	secondSlice := make([]int, 0)

	for _, line := range lines {
		nums := strings.Fields(line)
		firstNum, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatalf("Failed to proccess first number: %v", err)
		}

		secondNum, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatalf("Failed to proccess second number: %v", err)
		}

		firstSlice = append(firstSlice, firstNum)
		secondSlice = append(secondSlice, secondNum)
	}
	sort.Ints(firstSlice)
	sort.Ints(secondSlice)
	return firstSlice, secondSlice
}

func getDistances(firstSlice []int, secondSlice []int) int {
	distance := 0

	for index := range len(firstSlice) {
		distance += int(math.Abs(float64(firstSlice[index]) - float64(secondSlice[index])))
	}
	return distance
}

func getSimilarityScore(firstSlice []int, secondSlice []int) int {
	score := 0
	freq := map[int]int{}

	for _, num := range secondSlice {
		freq[num]++
	}

	for _, num := range firstSlice {
		score += freq[num] * num
	}
	return score
}

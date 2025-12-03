package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("Day 3, 2025!")
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to open file! %v", err)
	}
	banks := getBanks(string(file))
	// fmt.Println(banks)

	solutionOne := getTotalJoltage(banks)
	fmt.Printf("Solution one: %d", solutionOne)

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

func getTotalJoltage(banks [][]int) int {
	joltage := 0
	for _, bank := range banks {
		joltage += getLargestJoltage(bank)
	}
	return joltage
}
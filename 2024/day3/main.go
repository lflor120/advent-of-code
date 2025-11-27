package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"strconv"
)

func main() {
	fileContents, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file")
	}

	validMuls := getAllMuls(string(fileContents))
	fmt.Println(validMuls)
}

func getAllMuls(memory string) []string {
	regex := `mul\(\d{0,3},\d{0,3}\)`
	re, err := regexp.Compile(regex)
	if err != nil {
		log.Fatalf("Failed to compile regex: %v", err)
	}
	muls := re.FindAllString(memory, -1)
	return muls
}

func getMulOperands(mul string) (int, int) {
	regex := `\d{0,3},\d{0,3}`
	re, err := regexp.Compile(regex)
	if err != nil {
		log.Fatalf("Failed to compile operand regex: %v", err)
	}
	mulNums := re.FindAllString(mul, 1)[0]
	nums := strings.Split(mulNums, ",")
	
	numOne, err := strconv.Atoi(nums[0])
	if err != nil {
		log.Fatalf("Failed to convert %v to int", nums[0])
	}

	numTwo, err := strconv.Atoi(nums[1])
	if err != nil {
		log.Fatalf("Failed to convert %v to int", nums[1])
	}

	return numOne, numTwo
}
package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	fileContents, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file")
	}

	validMuls := getAllMuls(string(fileContents))
	fmt.Println(validMuls)
	// fmt.Printf("Solution One: %d\n", answerOne)
	
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

func parseOutOperands(mul string) (int, int) {


	return 0,0
}
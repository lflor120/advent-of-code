package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

// idea: not too hard, just gotta do some matrix manipulation stuff
// store a map like this map[col] = operation (*, +, etc)
// can use table driven dev to handle operations
// BETTER idea: rotate the string matrix and get the operation from the end
// after that just go row by row calculating
var operations = map[string]func([]int) int {
	"+": func(slice []int) int {
		total := 0
		for _, number := range slice {
			total += number
		}
		return total
	},
	"*": func(slice []int) int {
		total := 1
		for _, number := range slice {
			total *= number

		}
		return total
	},
}

func main() {
	fmt.Println("Day 6, 2025!")
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file! err: %v", err)
	}

	matrix := getStringMatrix(string(file))
	transposedMatrix := getTransposeMatrix(matrix)
	fmt.Println(transposedMatrix)
	solutionOne := Calculate(transposedMatrix)
	fmt.Printf("Solution one: %d\n", solutionOne)

}

// has to be a string since we have both numbers and symbols
func getStringMatrix(fileContents string) [][]string {
	matrix := make([][]string, 0)

	for _, line := range strings.Split(fileContents, "\n") {
		row := make([]string, 0)
		for _, num := range strings.Fields(line) {
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func getTransposeMatrix(matrix [][]string) [][]string {
	maxRow := len(matrix)
	maxCol := len(matrix[0])

	transposedMatrix := make([][]string, 0)
	for row := range maxCol {
		transposedRow := make([]string, 0)
		for col :=range maxRow {
			transposedRow = append(transposedRow, matrix[col][row])
		}
		transposedMatrix = append(transposedMatrix, transposedRow)
	}
	return transposedMatrix
}

func getEquation(matrixRow []string) ([]int, string) {
	rowIntSlice := make([]int, 0)
	for _, row := range matrixRow[:len(matrixRow) - 1] {
		rowInt, err := strconv.Atoi(row)
		if err != nil {
			log.Fatalf("Failed to convert %v to int!", rowInt)
		}
		rowIntSlice = append(rowIntSlice, rowInt)
	}

	operation := matrixRow[len(matrixRow)-1]
	return rowIntSlice, operation
}

// matrix should be transposed already
func Calculate(matrix [][]string) int {
	total := 0
	for _, row := range matrix {
		rowSlice, operation := getEquation(row)
		result := operations[operation](rowSlice)
		fmt.Printf("Current answer: %d\n", result)
		total += result
	}
	return total
}
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Direction struct {
	row, col int
}
var directions = []Direction{
	{row: 1, col: 0}, // down
	{row: -1, col: 0}, // up
	{row: 0, col: -1}, // left
	{row: 0, col: 1}, // right
	{row: 1, col: -1}, // down left
	{row: -1, col: -1}, // up left
	{row: 1, col: 1}, // down right
	{row: -1, col: 1}, // up right
}


func main() {
	fmt.Println("Day 4, 2025!")
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file! %v", err)
	}

	grid := createGrid(string(file))
	solutionOne := getTotalRolls(grid)
	fmt.Printf("Solution one: %d\n", solutionOne)

	solutionTwo := simulateRemoval(grid)
	fmt.Printf("Solution two: %d", solutionTwo)

}

func createGrid(fileContents string) [][]rune {
	grid := make([][]rune, 0)

	for _, line := range strings.Split(fileContents, "\n") {
		row := make([]rune, 0)
		for _, c := range line {
			row = append(row, c)
		}
		grid = append(grid, row)
	}
	return grid
}

func getTotalRolls(grid [][]rune) int {
	total := 0
	rowMax := len(grid)
	colMax := len(grid[0])

	for row := range rowMax {
		for col := range colMax {
			if grid[row][col] != '@' {
				continue
			}

			if neighbors := getNeighbors(row, col, grid); !isAccessible(neighbors) {
				continue
			}

			total++
		}
	}
	return total
}

func getTotalRollsAndRemove(grid [][]rune) ([][]rune, int) {
	total := 0
	rowMax := len(grid)
	colMax := len(grid[0])

	for row := range rowMax {
		for col := range colMax {
			if grid[row][col] != '@' {
				continue
			}

			if neighbors := getNeighbors(row, col, grid); !isAccessible(neighbors) {
				continue
			}
			grid[row][col] = '.'
			total++
		}
	}
	return grid, total
}


func simulateRemoval(grid [][]rune) int {
	totalRolls := 0
	simGrid := grid
	currTotal := 0
	for {
		simGrid, currTotal = getTotalRollsAndRemove(simGrid)

		if currTotal == 0 {
			break
		}

		totalRolls += currTotal
	}
	return totalRolls
}


func isAccessible(neighbors []rune) bool {
	rolls := 0
	for _, c := range neighbors {
		if c == '@' {
			rolls++
		}
	}
	return rolls < 4
}

// the meat of this problem
// use row, col values to find possible neighbors
// ensure values are in range of the grid
func getNeighbors(row, col int, grid [][]rune) []rune {
	neighbors := make([]rune, 0)
	for _, direction := range directions {
		newRow := row + direction.row
		newCol := col + direction.col

		if !inGridRange(newRow, newCol, grid) {
			continue
		}

		neighbors = append(neighbors, grid[newRow][newCol])
	}

	return neighbors
}

func inGridRange(row, col int, grid [][]rune) bool {
	maxRow := len(grid)
	maxCol := len(grid[0])

	if row < 0 || row >= maxRow {
		return false
	}
	if col < 0 || col >= maxCol {
		return false
	}
	return true
}
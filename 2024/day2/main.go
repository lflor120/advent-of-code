package main
 import (
	"fmt"
	"os"
 )

func main() {
	fmt.Println("Day 2!")


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
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sort"
)

func main() {
	fmt.Println("Day 5, 2025!")
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed reading file with error: %v", err)
	}
	fileContents := string(file)

	intervals := getIntervals(fileContents)
	ids := getIds(fileContents)

	solutionOne := getFresh(ids, intervals)
	fmt.Printf("Solution one: %d\n", solutionOne)

	solutionTwo := getAllFresh(intervals)
	fmt.Printf("Solution two: %d\n", solutionTwo)
}

func getIntervals(fileContents string) [][]int {
	intervals := make([][]int, 0)

	fileInterval := strings.Split(fileContents, "\n\n")[0]
	for _, intervalStr := range strings.Split(fileInterval, "\n") {
		interval := strings.Split(intervalStr, "-")
		start, err := strconv.Atoi(interval[0])
		if err != nil {
			log.Fatalf("Failed to convert %v to int!", interval[0])
		}

		stop, err := strconv.Atoi(interval[1])
		if err != nil {
			log.Fatalf("Failed to convert %v to int!", interval[1])
		}

		intervals = append(intervals, []int{start, stop})
	}
	return intervals
}

func getIds(fileContents string) []int {
	ids := make([]int, 0)

	fileIds := strings.Split(fileContents, "\n\n")[1]
	for _, id := range strings.Split(fileIds, "\n") {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalf("Failed to convert %v t int!", id)
		}
		ids = append(ids, idInt)
	}
	return ids
}

func mergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func (i, j int) bool {
		return intervals[i][0] < intervals[j][0] 
	})

	merged := make([][]int, 0)
	merged = append(merged, intervals[0])

	for _, interval := range intervals[1:] {
		n := len(merged)
		start, stop := interval[0], interval[1]
		prevStop := merged[n - 1][1]

		if start <= prevStop {
			merged[n - 1][1] = max(prevStop, stop)
			continue
		}

		merged = append(merged, interval)
	}
	fmt.Println("Slices sorted!")
	mergedDiff := len(intervals) - len(merged)
	fmt.Println("Original count: ", len(intervals))
	fmt.Println("Merged: ", mergedDiff)
	return merged
}

func getFresh(ids []int, intervals [][]int) int {
	fresh := 0

	mergedIntervals := mergeIntervals(intervals)

	fmt.Println("Counting fresh!")
	for _, id := range ids {
		if !isInRange(id, mergedIntervals) {
			continue
		}
		fresh++
	}

	return fresh
}

func isInRange(id int, intervals[][]int) bool {
	for _, interval := range intervals {
		start, stop := interval[0], interval[1]
		if id >= start && id <= stop {
			return true
		}
	}
	return false
}


//part 2
func getAllFresh(intervals [][]int) int {
	fresh := 0
	mergedIntervals := mergeIntervals(intervals)
	for _, interval := range mergedIntervals {
		start, stop := interval[0], interval[1]
		ids := (stop - start) + 1
		fmt.Printf("%d ids between %d - %d\n", ids, start, stop)
		fresh += ids
	}

	return fresh
}
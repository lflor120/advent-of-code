package main

import (
	"testing"
)

func TestLeftMax(t *testing.T) {
	test := []int{1, 4, 5, 2, 3, 5, 1}
	val, index := getLeftMostMax(test)

	if val != 5 && index != 2 {
		t.Fatal("leftMostMax failed")
	}
}

type testCase struct {
	bank   []int
	answer int64
	top    int
}

var tests = []testCase{
	{bank: []int{2, 3, 4, 5, 1}, answer: 451, top: 3},
	{bank: []int{2, 3, 4, 5, 1}, answer: 51, top: 2},
	{bank: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}, answer: 987654321111, top: 12},
	{bank: []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}, answer: 811111111119, top: 12},
	{bank: []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}, answer: 434234234278, top: 12}, // panic
	{bank: []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}, answer: 888911112111, top: 12},
}

func TestGetLargest12Joltage(t *testing.T) {
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			result := getLargestNJoltage(tt.bank, tt.top)
			if result != tt.answer {
				t.Errorf("getLargest12Joltage failed. Expected: %d | Got: %d", tt.answer, result)
			}
		})
	}
}

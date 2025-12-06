package main

import (
	"fmt"
	"log"
	"slices"
	"testing"
	"reflect"
)

type EquationTest struct {
	slice []string
	operation string
	expectedSlice []int
	expectedOp string
}

type MatrixTest struct {
	name string
	matrix, expected [][]string
}

type OperationTest struct {
	name, operation string
	slice []int
	expected int
}

func TestTransposeMatrix(t *testing.T) {
	cases := []MatrixTest{
		{
			name: "simple example",
			matrix: [][]string{
				{"1", "2"},
				{"3", "4"},
				{"+", "*"},
			},
			expected: [][]string{
				{"1", "3", "+"},
				{"2", "4", "*"},
			},
		},
		{
			name: "AoC example",
			matrix: [][]string{
				{"123", "328", "51", "64"},
				{"45", "64", "387", "23"},
				{"6", "98", "215", "314"},
				{"*", "+", "*", "+"},
			},
			expected: [][]string{
				{"123", "45", "6", "*"},
				{"328", "64", "98", "+"},
				{"51", "387", "215", "*"},
				{"64", "23", "314", "+"},
			},
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("name: %s", tc.name), func(t *testing.T) {
			result := getTransposeMatrix(tc.matrix)
			if !reflect.DeepEqual(result, tc.expected) {
				log.Fatalf("Test: %s failed!", tc.name)
			}
		})
	}

}


func TestGetEquation(t *testing.T) {
	cases := []EquationTest{
		{
			slice: []string{"123", "45", "6", "*"}, 
			operation: "*",
			expectedSlice: []int{123,45,6},
			expectedOp: "*",
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v using operation %s", tc.slice, tc.operation), func(t *testing.T) {
			resultSlice, resultOp := getEquation(tc.slice)
			if !slices.Equal(resultSlice, tc.expectedSlice) {
				log.Fatalf("Got: %v Expected: %v\n", resultSlice, tc.expectedSlice)
			}

			if resultOp != tc.expectedOp {
				log.Fatalf("Got: %v Expected: %v\n", resultOp, tc.expectedOp)
			}
		})
	}
}

func TestOperations(t *testing.T) {
	cases := []OperationTest{
		{
			name: "Aoc Example Multiply",
			operation: "*",
			slice: []int{123, 45, 6},
			expected: 33210,
		},
		{
			name: "Aoc Example Plus",
			operation: "+",
			slice: []int{328, 64, 98},
			expected: 490,
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("Test case: %v", tc.name), func(t *testing.T) {
			result := operations[tc.operation](tc.slice)
			if result != tc.expected {
				log.Fatalf("Operations failed for test: %v\nExpected: %d Got %d\n",tc.name, tc.expected, result)
			}
		})
	}
}
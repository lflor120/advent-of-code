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
				log.Fatalf("Got: %v Expected: %v", resultSlice, tc.expectedSlice)
			}

			if resultOp != tc.expectedOp {
				log.Fatalf("Got: %v Expected: %v", resultOp, tc.expectedOp)
			}
		})
	}
}
package main

import (
	// "strconv"
	"testing"
)

func TestIsRepeatedHappy(t *testing.T) {
	test := 123123
	result := isRepeatedTwice(test)

	if result != true {
		t.Error("isRepeated failed happy path")
	}
}

func TestIsRepeatedEvenLength(t *testing.T) {
	test := 11
	result := isRepeatedTwice(test)

	if result != true {
		t.Error("isRepeated failed for even length string")
	}
}

func TestIsRepeatedEvenLength2(t *testing.T) {
	test := 1010
	result := isRepeatedTwice(test)

	if result != true {
		t.Error("isRepeated failed for even length string")
	}
}

func TestIsRepeatedMultipleHappy(t *testing.T) {
	test := 1212121212
	result := isRepeatedMultiple(test)

	if result != true {
		t.Error("isRepeatedMultiple failed")
	}
}

type testCase struct {
	id int
}

var tests = []testCase{
	// {1},
	{11},
	{111},
	{1010},
	{1212121212},
	{1188511885},
	{222222},
	{446446},
	{38593859},
	{565656},
	{824824824},
	{2121212121},
	{181818},
	{118811881188}, 
}

func TestIsInvalid(t *testing.T) {
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			result := isInvalid(tt.id)
			if result != true {
				t.Errorf("isInvalid failed for ID: %d", tt.id)
			}
		})
	}
}

func TestRemoveAdjacent(t *testing.T) {
	test := "112233112233"
	result := removeAdjacent(test)

	if result != "123123" {
		t.Error(result)
	}

}

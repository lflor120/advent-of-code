package main

import "testing"

func TestIsRepeatedHappy(t *testing.T) {
	test := 123123
	result := isRepeatedNumber(test)

	if result != true {
		t.Error("isRepeated failed happy path")
	}
}

func TestIsRepeatedEvenLength(t *testing.T) {
	test := 11
	result := isRepeatedNumber(test)

	if result != true {
		t.Error("isRepeated failed for even length string")
	}
}

func TestIsRepeatedEvenLength2(t *testing.T) {
	test := 1010
	result := isRepeatedNumber(test)

	if result != true {
		t.Error("isRepeated failed for even length string")
	}
}
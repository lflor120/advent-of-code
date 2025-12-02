package main

import "testing"

func TestIsRepeatedHappy(t *testing.T) {
	test := 123123
	result := isRepeatedNumber(test)

	if result != true {
		t.Error("isRepeated failed happy path")
	}
}
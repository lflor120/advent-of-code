package main

import (
	"testing"
	"slices"
)

func TestGetMuls(t *testing.T) {
	memory := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expected := []string{
		"mul(2,4)",
		"mul(5,5)",
		"mul(11,8)",
		"mul(8,5)",
	}
	muls := getAllMuls(memory)

	if len(muls) != 4 {
		t.Errorf("Expected 4 muls, but got %d", len(muls))
	}

	if !slices.Equal(expected, muls) {
		t.Errorf("Mul output is incorrect\nExpected: %v\nResult: %v", expected, muls)
	}

}

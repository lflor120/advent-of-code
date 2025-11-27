package main

import (
	"log"
	"slices"
	"testing"
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

func TestParseOperands(t *testing.T) {
	mul := "mul(123,65)"
	resultOne, resultTwo := getMulOperands(mul)

	if resultOne != 123 {
		log.Fatalf("getMulOperands was expected: 123 but got %v", resultOne)
	}

	if resultTwo != 65 {
		log.Fatalf("getMulOperands was expected: 65 but got %v", resultTwo)
	}

}

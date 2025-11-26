package main

import "testing"

func TestIncreasingTrue(t *testing.T) {
	test := []int{1,2,3,4,5}
	result := isIncreasing(test)

	if result != true {
		t.Error("increasing failed happy path")
	}
}

func TestIncreasingFalse(t *testing.T) {
	test := []int{1,2,3,1,6}
	result := isIncreasing(test)

	if result != false {
		t.Error("increasing failed sad path")
	}
}

func TestDecreasingTrue(t *testing.T) {
	test := []int{6,5,4,3,2,1}
	result := isDecreasing(test)

	if result != true {
		t.Error("decreasing failed happy path")
	}
}

func TestDecreasingFalse(t *testing.T) {
	test := []int{6,5,4,3,4,1}
	result := isDecreasing(test)

	if result != false {
		t.Error("decreasing failed sad path")
	}
}

func TestDifferenceHappy(t *testing.T) {
	test := []int{1,3,5}
	result := validDifference(test, 0)

	if result != true {
		t.Error("valid difference failed happy path")
	}
}

func TestDifferenceSad(t *testing.T) {
	test := []int{1,5,10}
	result := validDifference(test, 0)

	if result != false {
		t.Error("valid difference failed sad path")
	}
}
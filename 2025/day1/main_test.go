package main

import "testing"

func TestCircularLL(t *testing.T) {
	ll := CircularLL{}
	ll.Setup(3)

	order := []int{0,1,2,0,1,2}

	trav := ll.current
	for i := range 6 {
		if trav.value != order[i] {
			t.Errorf("Circular LL not working. Expected %d, but got %d", order[i], trav.value)
		}
		trav = trav.next
	}
}

func TestCircularLLBackwards(t *testing.T) {
	ll := CircularLL{}
	ll.Setup(3)

	order := []int{0,1,2,0,1,2}

	trav := ll.current.prev
	for i := len(order) - 1; i >= 0; i-- {
		if trav.value != order[i] {
			t.Errorf("Circular LL not working backwards. Expected %d, but got %d", order[i], trav.value)
		}
		trav = trav.prev
	}
}

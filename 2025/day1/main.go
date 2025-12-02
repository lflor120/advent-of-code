package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
)

// thought process, create a circular doubly linked list?
// not most optimal but easiest to read and possibly extend 🤔

func main() {
	fmt.Println("Day 1 2025!")
	padlock := CircularLL{}
	padlock.Setup(100)

	// read input file
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file with err: %v", err)
	}

	instructions := parseInstructions(string(file))
	solutionOne := padlock.getPasscode(instructions)
	fmt.Printf("Solution one: %d\n", solutionOne)
}

func (c *CircularLL) getPasscode(instructions []string) int {
	zeros := 0
	// setup current at 0
	for range 50 {
		c.current = c.current.next
	}
	
	// process instructions
	for _, instruction := range instructions {
		direction := instruction[0]
		iterations, err := strconv.Atoi(instruction[1:])
		if err != nil {
			log.Fatalf("Failed to convert: %v to int!", instruction[1:])
		}

		for range iterations {
			switch direction {
				case 'L':
					c.current = c.current.prev
				case 'R':
					c.current = c.current.next
			}
			if c.current.value == 0 {
				zeros++
			}
		}
		// if c.current.value == 0 {
		// 	zeros++
		// }
	}
	return zeros
}

func parseInstructions(fileContents string) []string {
	lines := strings.Split(fileContents, "\n")
	return lines
}

type Node struct {
	next, prev *Node
	value int
}

type CircularLL struct {
	current, head, tail *Node
}

func (c *CircularLL) Insert(value int) {
	newNode := Node{value: value}
	if c.head == nil {
		// set up initial
		c.head = &newNode
		c.tail = &newNode
		c.current = c.head
		c.head.prev = c.tail
		c.tail.next = c.head
		return
	}

	// set up existing LL
	curr := c.tail
	curr.next = &newNode 
	newNode.prev = curr
	c.tail = &newNode

	// circular binding
	c.tail.next = c.head
	c.head.prev = c.tail
}

func (c *CircularLL) Setup(size int) {
	for i := range size {
		c.Insert(i)
	}
}
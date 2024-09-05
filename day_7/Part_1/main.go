package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var wires map[string]uint16

func main() {
	wires = make(map[string]uint16)

	// Read and parse the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	instructions := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		instructions[parts[1]] = parts[0] // Store the instruction for each wire
	}

	// Evaluate wire 'a'
	result := evaluate("a", instructions)
	fmt.Println("Signal on wire a:", result)
}

func evaluate(wire string, instructions map[string]string) uint16 {
	// Check if the value is already computed
	if value, exists := wires[wire]; exists {
		return value
	}

	instruction := instructions[wire]

	// Check if the instruction is a constant value (number)
	if num, err := strconv.Atoi(instruction); err == nil {
		wires[wire] = uint16(num)
		return uint16(num)
	}

	// Handle different types of operations
	parts := strings.Split(instruction, " ")

	var result uint16
	switch len(parts) {
	case 1: // Direct assignment
		result = evaluate(parts[0], instructions)
	case 2: // NOT operation
		operand := evaluate(parts[1], instructions)
		result = ^operand
	case 3: // AND, OR, LSHIFT, RSHIFT
		left := getValue(parts[0], instructions)
		right := getValue(parts[2], instructions)

		switch parts[1] {
		case "AND":
			result = left & right
		case "OR":
			result = left | right
		case "LSHIFT":
			result = left << right
		case "RSHIFT":
			result = left >> right
		}
	}

	// Cache the result
	wires[wire] = result
	return result
}

// getValue retrieves the value either from a constant or by evaluating a wire
func getValue(operand string, instructions map[string]string) uint16 {
	if num, err := strconv.Atoi(operand); err == nil {
		return uint16(num) // If it's a number, return it directly
	}
	return evaluate(operand, instructions) // Otherwise, evaluate the wire
}

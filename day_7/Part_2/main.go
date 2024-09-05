package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	wires := make(map[string]uint16)

	// Function to reset the wire map
	resetWires := func() {
		for k := range wires {
			delete(wires, k)
		}
	}

	// Function to evaluate instructions and populate the wire values
	evaluateWires := func(instructions []string) {
		for len(instructions) > 0 {
			remainingInstructions := []string{}

			for _, instruction := range instructions {
				parts := strings.Split(instruction, " -> ")
				lhs := parts[0]
				wire := parts[1]

				if _, exists := wires[wire]; exists {
					continue
				}

				lhsParts := strings.Split(lhs, " ")
				var result uint16

				if len(lhsParts) == 1 {
					if val, err := strconv.Atoi(lhsParts[0]); err == nil {
						result = uint16(val)
					} else if val, exists := wires[lhsParts[0]]; exists {
						result = val
					} else {
						remainingInstructions = append(remainingInstructions, instruction)
						continue
					}
				} else if len(lhsParts) == 2 {
					if val, exists := wires[lhsParts[1]]; exists {
						result = ^val
					} else {
						remainingInstructions = append(remainingInstructions, instruction)
						continue
					}
				} else if len(lhsParts) == 3 {
					left, right := lhsParts[0], lhsParts[2]
					var leftVal, rightVal uint16

					if val, err := strconv.Atoi(left); err == nil {
						leftVal = uint16(val)
					} else if val, exists := wires[left]; exists {
						leftVal = val
					} else {
						remainingInstructions = append(remainingInstructions, instruction)
						continue
					}

					if val, err := strconv.Atoi(right); err == nil {
						rightVal = uint16(val)
					} else if val, exists := wires[right]; exists {
						rightVal = val
					} else {
						remainingInstructions = append(remainingInstructions, instruction)
						continue
					}

					switch lhsParts[1] {
					case "AND":
						result = leftVal & rightVal
					case "OR":
						result = leftVal | rightVal
					case "LSHIFT":
						result = leftVal << rightVal
					case "RSHIFT":
						result = leftVal >> rightVal
					}
				}

				wires[wire] = result
			}

			instructions = remainingInstructions
		}
	}

	// Read and parse the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Store instructions
	var instructions []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	// Step 1: Evaluate wires normally to get signalA for wire 'a'
	evaluateWires(instructions)
	signalA := wires["a"]
	fmt.Println("Signal on wire a in first run:", signalA)

	// Step 2: Override wire 'b' with signalA
	resetWires() // Reset all wires
	for i, instruction := range instructions {
		if strings.HasSuffix(instruction, " -> b") {
			instructions[i] = fmt.Sprintf("%d -> b", signalA)
			break
		}
	}

	// Step 3: Re-evaluate wires with the modified instructions
	evaluateWires(instructions)
	fmt.Println("Signal on wire a after overriding b:", wires["a"])
}

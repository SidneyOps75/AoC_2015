package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Create a 1000x1000 grid initialized to false (lights off)
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var action string
		if strings.HasPrefix(line, "turn on") {
			action = "turn on"
			line = strings.TrimPrefix(line, "turn on ")
		} else if strings.HasPrefix(line, "turn off") {
			action = "turn off"
			line = strings.TrimPrefix(line, "turn off ")
		} else if strings.HasPrefix(line, "toggle") {
			action = "toggle"
			line = strings.TrimPrefix(line, "toggle ")
		} else {
			continue
		}

		// Split the coordinates part into two parts: start and end
		parts := strings.Split(line, " through ")
		if len(parts) != 2 {
			fmt.Println("Invalid input line:", line)
			continue
		}

		startCoords := strings.Split(parts[0], ",")
		endCoords := strings.Split(parts[1], ",")

		if len(startCoords) != 2 || len(endCoords) != 2 {
			fmt.Println("Invalid coordinates in line:", line)
			continue
		}

		// Convert the coordinates from strings to integers
		x1, err1 := strconv.Atoi(startCoords[0])
		y1, err2 := strconv.Atoi(startCoords[1])
		x2, err3 := strconv.Atoi(endCoords[0])
		y2, err4 := strconv.Atoi(endCoords[1])

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			fmt.Println("Error converting coordinates in line:", line)
			continue
		}

		// Apply the instruction to the grid
		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				switch action {
				case "turn on":

					grid[i][j]++
				case "turn off":
					if grid[i][j] > 0 {
						grid[i][j]--
					}

				case "toggle":
					grid[i][j] += 2
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Count the number of lights that are on
	totalbrightness := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			totalbrightness += grid[i][j]
		}
	}

	fmt.Println("TotalBrightness:", totalbrightness)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	x, y int
}

func containsCoordinate(slice []Coordinate, coord Coordinate) bool {
	for _, v := range slice {
		if v == coord {
			return true
		}
	}
	return false
}

func countHousesWithPresents(directions string) int {
	x, y := 0, 0
	visited := []Coordinate{{x, y}}

	for _, direction := range directions {
		switch direction {
		case '>':
			x++
		case '<':
			x--
		case '^':
			y++
		case 'v':
			y--
		}

		currentCoord := Coordinate{x, y}
		if !containsCoordinate(visited, currentCoord) {
			visited = append(visited, currentCoord)
		}
	}

	return len(visited)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	directions := scanner.Text()

	housesWithPresents := countHousesWithPresents(directions)
	fmt.Printf("Number of houses that receive at least one present: %d\n", housesWithPresents)
}

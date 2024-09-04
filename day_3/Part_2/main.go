package main

import (
	"bufio"
	"fmt"
	"os"
)

const gridSize = 1000 // Adjust this based on expected grid size

type Point struct {
	x, y int
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

	// Part 1: Original Santa
	housesVisited := countHousesVisited(directions, false)
	fmt.Printf("Part 1: Number of houses that received at least one present: %d\n", housesVisited)

	// Part 2: Santa and Robo-Santa
	housesVisited = countHousesVisited(directions, true)
	fmt.Printf("Part 2: Number of houses that received at least one present: %d\n", housesVisited)
}

func countHousesVisited(directions string, useRoboSanta bool) int {
	grid := make([][]bool, gridSize)
	for i := range grid {
		grid[i] = make([]bool, gridSize)
	}

	santa := Point{gridSize / 2, gridSize / 2}
	roboSanta := Point{gridSize / 2, gridSize / 2}
	grid[santa.y][santa.x] = true
	housesVisited := 1

	for i, move := range directions {
		var current *Point
		if useRoboSanta && i%2 == 1 {
			current = &roboSanta
		} else {
			current = &santa
		}

		switch move {
		case '^':
			current.y--
		case 'v':
			current.y++
		case '>':
			current.x++
		case '<':
			current.x--
		}

		if !grid[current.y][current.x] {
			grid[current.y][current.x] = true
			housesVisited++
		}
	}

	return housesVisited
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	openFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf("Error opening file:%s\n", err)
		return
	}

	defer openFile.Close()

	scanner := bufio.NewScanner(openFile)
	requiredArea := 0
	for scanner.Scan() {
		dimensions := scanner.Text()
		length, width, height := GetDimensions(dimensions)
		side1, side2, side3 := (length * width), (width * height), (height * length)
		dimensionAreas := []int{side1, side2, side3}
		sort.Ints(dimensionAreas)
		extrapaper := dimensionAreas[0]
		surfaceArea := (2 * length * width) + (2 * width * height) + (2 * height * length)
		totalArea := extrapaper + surfaceArea

		fmt.Printf("Total square feet area: %v\n", totalArea)

		requiredArea += totalArea

		fmt.Printf("Total square feet required: %v\n\n", requiredArea)

	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file:%s\n", err)
	}

}

func GetDimensions(dimensions string) (int, int, int) {
	parts := strings.Split(dimensions, "x")
	length, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Printf("Error parsing length from dimensions: %s\n", dimensions)
		return 0, 0, 0
	}
	width, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Printf("Error parsing width from dimensions: %s\n", dimensions)
		return 0, 0, 0
	}
	height, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Printf("Error parsing height from dimensions: %s\n", dimensions)
		return 0, 0, 0
	}
	return length, width, height
}

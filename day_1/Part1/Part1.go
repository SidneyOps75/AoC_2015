package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	OpenFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer OpenFile.Close()

	scanner := bufio.NewScanner(OpenFile)
	for scanner.Scan() {
		characters := scanner.Text()
		floor := 0
		for _, ch := range characters {
			if ch == '(' {
				floor++
			} else if ch == ')' {
				floor--
			}
		}
		fmt.Println(floor)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}
}

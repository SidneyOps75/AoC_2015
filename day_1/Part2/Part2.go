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
		for indx, ch := range characters {
			if ch == '(' {
				floor++
			} else if ch == ')' {
				floor--
			}
			if floor == -1 {
				fmt.Printf("The first instance of basement: %d\n", (indx + 1))
				break
			}
		}
		fmt.Printf("The current floor number is: %d\n", floor)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	OpenFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("file not found")
	}

	defer OpenFile.Close()

	scanner := bufio.NewScanner(OpenFile)

	for scanner.Scan() {
		characters := scanner.Text()

	}

}

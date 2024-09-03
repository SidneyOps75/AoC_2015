package main

import (
	"bufio"
	"fmt"
)

func main() {
	OpenFile , err := os.Open("./input")
	scanner, err := bufio.NewScanner("./input.txt")
	if err != nil {
		fmt.Println("file not found")
	}
	scanner.Scan(){
		text, err := scanner.Text()
		if err != nil {
			fmt.Println("failed to scan the text")
		}
	}

}

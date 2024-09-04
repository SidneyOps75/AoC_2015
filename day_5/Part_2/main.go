package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IsNice(s string) bool {
	// first condition : Contains a pair of any letters that appears at least twice in the string without overlapping
	var hasPair bool
	for i := 0; i < len(s)-1; i++ {
		pair := s[i:(i + 2)]
		if strings.Count(s, pair) > 1 {
			hasPair = true
			break
		}
	}
	if !hasPair {
		return false
	}

	// Second Condition ; Contains atleast one letter which repeats exactly one letter between them
	var hasRepeat bool
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			hasRepeat = true
			break
		}
	}
	if !hasRepeat {
		return false
	}

	return true
}

func main() {
	OpenFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}

	defer OpenFile.Close()

	scanner := bufio.NewScanner(OpenFile)

	niceCount := 0

	for scanner.Scan() {
		characters := scanner.Text()
		if IsNice(characters) {
			niceCount++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error Scanning File", err)
	}

	fmt.Printf("Number of nice strings :%v\n", niceCount)
}

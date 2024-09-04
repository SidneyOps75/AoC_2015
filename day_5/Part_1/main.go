package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IsNice(s string) bool {
	// First condition : It contains at least three vowels
	vowels := "aeiou"
	vowelCount := 0
	for _, ch := range s {
		if strings.ContainsRune(vowels, ch) {
			vowelCount++
		}
	}
	if vowelCount < 3 {
		return false
	}

	// Second condition : Contains atleast one letter that appears twice in a row
	var hasDouble bool
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			hasDouble = true
			break
		}
	}
	if !hasDouble {
		return false
	}

	// Third condition : It does not contain mentioned strings
	forbidden := []string{"ab", "cd", "pq", "xy"}
	for _, chars := range forbidden {
		if strings.Contains(s, chars) {
			return false
		}
	}
	return true
}

func main() {
	OpenFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("file not found")
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
		fmt.Println(err)
	}
	fmt.Printf("Nice Strings :%v\n", niceCount)
}

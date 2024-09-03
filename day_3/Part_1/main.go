package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	secretKey := "iwrupvqb"
	targetPrefix := "00000"
	var number int
	for {
		// Concatenate secretKey with the current number
		data := secretKey + Itoa(number)

		// Compute the MD5 hash
		hash := md5.Sum([]byte(data))

		// Convert hash to a hexadecimal string
		hashString := fmt.Sprintf("%x", hash)

		// Check if the hash starts with the target prefix
		if hashString[:5] == targetPrefix {
			fmt.Printf("Found the number: %d\n", number)
			break
		}

		// Increment the number
		number++
	}
}

func Itoa(s int) string {
	sign := ""
	if s == 0 {
		return "0"
	}
	if s < 0 {
		sign = "-"
		s = -s
	}
	var digits []rune
	for s > 0 {
		digit := s % 10
		digits = append([]rune{rune(digit + '0')}, digits...)
		s /= 10
	}
	return sign + string(digits)
}

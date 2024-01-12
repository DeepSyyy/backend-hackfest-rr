package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomClanTag() string {
	rand.Seed(time.Now().UnixNano())

	// Define the pattern for the clan tag
	pattern := "#XXXXXX"

	// Replace 'X' with random alphanumeric characters
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == 'X' {
			// ASCII values for alphanumeric characters (48-57: 0-9, 65-90: A-Z, 97-122: a-z)
			randomChar := rand.Intn(62)
			if randomChar < 10 {
				pattern = pattern[:i] + string(rune(randomChar+48)) + pattern[i+1:]
			} else if randomChar < 36 {
				pattern = pattern[:i] + string(rune(randomChar+55)) + pattern[i+1:]
			} else {
				pattern = pattern[:i] + string(rune(randomChar+61)) + pattern[i+1:]
			}
		}
	}

	return pattern
}

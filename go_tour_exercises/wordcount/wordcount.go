package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// Create an empty map to store the count of each word.
	result := make(map[string]int)

	// Split the string into words using strings.Fields.
	// This handles splitting on spaces and also cleans up any leading or trailing spaces.
	words := strings.Fields(s)

	// Iterate over each word.
	for _, word := range words {
		// Increase the count for each word in the map.
		// The map will automatically handle new words.

		// incase you want to count words ignoring the the case
		// word := strings.ToLower(word)

		result[word]++
	}

	// Return the map with word counts.
	return result
}

func main() {
	wc.Test(WordCount)
}

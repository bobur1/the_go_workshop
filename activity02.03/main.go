package main

import (
	"fmt"
)

func main() {
	var highestCount int
	var mostPopularWord string
	words := map[string]int{
		"Gonna": 3,
		"You": 3,
		"Give": 2,
		"Never": 1,
		"Up": 4,
		}

	for word, count := range words {
		if highestCount < count {
			highestCount = count
			mostPopularWord = word
		}
	}

	fmt.Printf("Most popular word: %s\n", mostPopularWord)
	fmt.Printf("With a count of  : %d", highestCount)
}

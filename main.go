package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	wordsFileName     = "english-words.txt"
	quartilesFileName = "quartiles.txt"
)

func main() {
	now := time.Now()

	allWords, err := parseFile(wordsFileName)
	if err != nil {
		os.Exit(1)
	}

	quartiles, err := parseFile(quartilesFileName)
	if err != nil {
		os.Exit(1)
	}

	wordSet := make(map[string]struct{}, len(allWords))
	for _, word := range allWords {
		wordSet[word] = struct{}{}
	}

	guesses := generateGuesses(quartiles)
	for _, guess := range guesses {
		if _, found := wordSet[guess]; found {
			fmt.Println("⚡️ found", guess)
		}
	}

	fmt.Println("Time:", time.Since(now))
}

func parseFile(fileName string) ([]string, error) {
	file, fileErr := os.Open(fileName)
	if fileErr != nil {
		fmt.Println("Error opening file:", fileErr)
		return nil, fileErr
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		fmt.Println("Error getting file stats:", statsErr)
		return nil, statsErr
	}

	bytes := make([]byte, stats.Size())
	_, readErr := file.Read(bytes)
	if readErr != nil {
		fmt.Println("Error reading file:", readErr)
		return nil, readErr
	}

	rawText := string(bytes)
	words := strings.Split(rawText, "\n")
	for i := range words {
		words[i] = strings.TrimSpace(words[i])
	}
	return words, nil
}

// Combine quartiles into all variations of 4-tile potential words
func generateGuesses(quartiles []string) []string {
	possibleWords := []string{}
	length := len(quartiles)

	for a := range length {
		for b := range length {
			for c := range length {
				for d := range length {
					if a != b && a != c && a != d && b != c && b != d && c != d {
						possibleWords = append(possibleWords, quartiles[a]+quartiles[b]+quartiles[c]+quartiles[d])
					}
				}
			}
		}
	}

	return possibleWords
}

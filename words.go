package main

import (
	"bufio"
	"math/rand"
	"os"
)

// Returns a hashset of all words in the given file
func loadWords(path string) (map[string]struct{}, error) {
	// Allocate set with empty struct
	wordlist := make(map[string]struct{})

	// retrieve the file from path
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close() // runs on any return case, including errors

	// Add each token to the set
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordlist[scanner.Text()] = struct{}{}
	}
	return wordlist, scanner.Err()
}

// Return array: we dont need lookup for a random pick?
func loadAnswers(path string) ([]string, error) {
	var wordlist []string

	// retrieve the file from path
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close() // runs on any return case, including errors

	// Add each token to the set
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}
	return wordlist, scanner.Err()
}

func newWord(wordlist []string) string {
	size := len(wordlist)
	// error case
	if size == 0 {
		return "ERROR"
	}
	return wordlist[rand.Intn(size)]
}

// Run on guess submission
func isMatching(guess []letter, answer string) ([]letter, bool) {
	word := guess
	matching := true

	// Cast answer to map for lookup
	inAnswer := make(map[rune]struct{})
	for _, char := range answer {
		inAnswer[rune(char)] = struct{}{}
	}

	//
	for i, char := range word {
		if char.char == rune(answer[i]) {
			// Correct VAL and POS
			char.state = Correct
		} else if _, ok := inAnswer[char.char]; ok {
			// Correct VAL
			char.state = Partial
			matching = false
		} else {
			// Not in answer
			char.state = Wrong
			matching = false
		}
	}
	return word, matching
}

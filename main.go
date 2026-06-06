package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

// Global lookup
var allowedWords map[string]struct{}
var realWords []string

func main() {
	var err error

	allowedWords, err = loadWords("allowed_wordlist.txt")
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	realWords, err = loadAnswers("real_wordlist.txt")
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	p := tea.NewProgram(initModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

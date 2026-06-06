package main

import (
	tea "charm.land/bubbletea/v2"
)

// Enum
type letterState int

const (
	Unset letterState = iota
	Wrong
	Partial
	Correct
)

type letter struct {
	char  rune
	state letterState
}

// ---------------- Model ----------------
type model struct {
	done    bool            // If session is over
	answer  string          // Correct string
	letters map[rune]letter // Which letters have been guessed
	guesses [][]letter      // Guessed words in session
	wordPos int             // Current position in guesses
	charPos int             // Currect position in the word
}

func resetLetters() map[rune]letter {
	letters := "qwertyuiopasdfghjklzxcvbnm"

	letterMap := make(map[rune]letter)

	for _, val := range letters {
		letterMap[val] = letter{
			state: Unset,
			char:  val,
		}
	}

	return letterMap
}

func resetModel() (tea.Model, tea.Cmd) {
	m := model{
		// m.resetLetters()
		// m.done = false
		// m.wordPos = 0
		// m.charPos = 0
		// m.answer = newWord(realWords)
		// m.guesses = [][]letter{}
		letters: resetLetters(),
		done:    false,
		wordPos: 0,
		charPos: 0,
		answer:  newWord(realWords),
		guesses: [][]letter{},
	}
	return m, nil
}

func initModel() tea.Model {
	m, _ := resetModel()
	return m
}

// Run when program starts
func (m model) Init() tea.Cmd {
	// return doTick()
	return nil
}

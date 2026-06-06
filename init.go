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
}

func (m *model) resetLetters() {
	letters := "qwertyuiopasdfghjklzxcvbnm"

	letterMap := make(map[rune]letter)

	for _, val := range letters {
		letterMap[val] = letter{
			state: Unset,
			char:  val,
		}
	}

	m.letters = letterMap
}

func (m *model) resetModel() {
	m.resetLetters()

	m.done = false
	m.wordPos = 0
	m.answer = newWord(realWords)
}

func initModel() model {
	m := model{}
	m.resetModel()
	return m
}

// Run when program starts
func (m model) Init() tea.Cmd {
	// return doTick()
	return nil
}

package main

import (
	tea "charm.land/bubbletea/v2"
)

// ENUM for tracking guessed letter states
type letterState int

const (
	Unset letterState = iota
	Wrong
	Partial
	Correct
)

// ENUM for tracking if the user has won or not
type gameState int

const (
	InProgress gameState = iota
	Won
	Lost
)

type letter struct {
	char  rune
	state letterState
}

// ---------------- Model ----------------
type model struct {
	gameState gameState       // If session is over
	answer    []letter        // Correct string
	letters   map[rune]letter // Which letters have been guessed
	guesses   [][]letter      // Guessed words in session
	wordPos   int             // Current position in guesses
	charPos   int             // Currect position in the word
	errorMsg  string          // When guess submission isnt valid
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
	m.gameState = InProgress
	m.resetLetters()
	m.answer = newWord(realWords)
	m.guesses = make([][]letter, 6)
	m.wordPos = 0
	m.charPos = 0
	m.errorMsg = ""
}

func initModel() tea.Model {
	var m model
	m.resetModel()
	return m
}

// Run when program starts
func (m model) Init() tea.Cmd {
	// return doTick()
	return nil
}

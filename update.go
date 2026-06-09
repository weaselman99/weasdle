package main

import (
	tea "charm.land/bubbletea/v2"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKey(msg)
		// case TickMsg: // Called every milisecond
	}
	return m, nil
}

// Handles any keypress event
func (m *model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	key := msg.Key()

	// Default keys - activated whenever during session
	switch {
	case key.Code == tea.KeyEsc: // ESC - Quit app
		return m, tea.Quit
	case key.Code == tea.KeyTab: // TAB - reset model with new answer
		m.resetModel()
		return m, nil
	}

	// Check whether game is in progress
	if m.gameState != InProgress {
		return m, nil
	}

	// Keys only allowed when user is currently playing
	switch {
	case key.Code == tea.KeyEnter: // ENTER - submit current guess
		return m.handleEnter()
	case key.Code == tea.KeyBackspace: // BACKSPACE - delete last letter in current guess
		return m.handleBackspace()
	case len(key.Text) == 1: // Character key press
		return m.handleRune(msg)
	}

	return m, nil
}

// Handles backspc - remove last letter
func (m *model) handleBackspace() (tea.Model, tea.Cmd) {
	// Empty word
	if m.charPos == 0 {
		return m, nil
	}

	m.guesses[m.wordPos] = m.guesses[m.wordPos][:m.charPos-1]

	m.charPos--
	return m, nil
}

// Handles enter - submission of word
func (m *model) handleEnter() (tea.Model, tea.Cmd) {
	word := m.guesses[m.wordPos]
	// Check if word is complete
	if len(word) < 5 {
		return m, nil
	}

	// Check guess and change states
	var matching bool
	word, matching = isMatching(word, m.answer)
	m.guesses[m.wordPos] = word

	if matching {
		// If guess is correct
		m.gameState = Won
		return m, nil
	}

	if _, ok := allowedWords[letterToString(word)]; !ok {
		// If guess isnt valid
		m.errorMsg = "Not a valid word"
		return m, nil
	}

	// Advance wordpos and reset charpos
	m.wordPos++
	m.charPos = 0
	return m, nil
}

// Handles the default key press
func (m *model) handleRune(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	key := rune(msg.Key().Text[0]) // cast key to rune

	// Ensure it's a valid letter
	if _, ok := m.letters[key]; !ok {
		return m, nil // nothing happens
	}

	// Ensure it's in bounds
	if m.charPos >= 5 {
		return m, nil
	}

	// Create new word if current word doesn't exist yet
	if len(m.guesses) <= m.wordPos {
		m.guesses = append(m.guesses, []letter{})
	}

	// Append new char
	m.guesses[m.wordPos] = append(m.guesses[m.wordPos], letter{
		char:  key,
		state: Unset,
	})

	m.charPos += 1

	return m, nil
}

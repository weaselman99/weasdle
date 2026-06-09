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
	m.errorMsg = ""
	// Empty word
	if m.charPos == 0 {
		return m, nil
	}

	// Replace current word with []slice with last element removed
	m.guesses[m.wordPos][m.charPos-1] = letter{
		char:  ' ',
		state: Empty,
	}
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

	// If guess isnt valid, don't update anything
	if _, ok := allowedWords[letterToString(word)]; !ok {
		m.errorMsg = "Not a valid word"
		return m, nil
	}

	// Check guess and change states
	result, matching := isMatching(word, m.answer)

	m.guesses[m.wordPos] = result
	m.updateLetters(result)

	if matching {
		// If guess is correct
		m.gameState = Won
		return m, nil
	}

	// Advance wordpos and reset charpos
	m.wordPos++
	m.charPos = 0

	// Check whether this was the last guess
	if m.wordPos > 5 {
		m.gameState = Lost
		m.errorMsg = "You lost!"
		return m, nil
	}

	return m, nil
}

// Changes the letterStates in letters from a given guess
func (m *model) updateLetters(word []letter) {
	for _, char := range word {
		let := m.letters[char.char]
		if let.state == Correct {
			continue
		}
		let.state = char.state
		m.letters[char.char] = let
	}
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

	// Update letter
	m.guesses[m.wordPos][m.charPos] = letter{
		char:  key,
		state: Unset,
	}

	m.charPos += 1

	return m, nil
}

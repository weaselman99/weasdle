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
func (m model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	key := msg.Key()
	switch {
	case key.Code == tea.KeyEsc: // ESC - Quit app
		return m, tea.Quit
	case key.Code == tea.KeyTab: // TAB
		m.resetModel()
		return m, nil
	case len(key.Text) == 1: // Character key press
		if !m.done {
			return m.handleRune(msg)
		}
	}
	return m, nil
}

// Handles the default key press
func (m model) handleRune(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	key := msg.Key()

	// Ensure it's a valid letter
	if _, ok := m.letters[rune(key.Text[0])]; !ok {
		return m, nil // nothing happens
	}

	// kkk

	return m, nil
}

package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	// dELTE test
	var guesses string
	for _, word := range m.guesses {
		guesses += "\n"
		for _, char := range word {
			guesses += string(char.char)
		}
	}
	testString := fmt.Sprintf("%s", letterToString(m.answer)) + guesses

	// Init view
	v := tea.NewView(testString)
	v.AltScreen = true
	return v
}

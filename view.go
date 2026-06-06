package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	testString := fmt.Sprintf("%s", m.answer)
	v := tea.NewView(testString)
	v.AltScreen = true
	return v
}

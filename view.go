package main

import (
	"fmt"

	"charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// ---------------- STYLES	----------------
// Colours
var (
	// WHITE  = lipgloss.Color("#d8dee9")
	// BLACK  = lipgloss.Color("#4c566a")
	// GRAY   = lipgloss.Color("#565f89")
	// RED    = lipgloss.Color("#bf616a")
	// GREEN  = lipgloss.Color("#a3be8c")
	// YELLOW = lipgloss.Color("#ebcb8b")
	WHITE  = lipgloss.Color("#d8dee9")
	BLACK  = lipgloss.Black
	GRAY   = lipgloss.Color("#444444")
	RED    = lipgloss.Red
	GREEN  = lipgloss.Green
	YELLOW = lipgloss.Yellow
)

var textStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(WHITE)

var boxStyle = lipgloss.NewStyle().
	Inherit(textStyle).
	Margin(0, 1, 1).
	Padding(1, 3)

var rowStyle = lipgloss.NewStyle()

// ---------------- VIEWS	----------------

func (m model) renderWords() string {
	var result string
	for _, guess := range m.guesses {
		var word string
		for _, char := range guess {
			// Change BG based on the state
			let := boxStyle
			switch char.state {
			case Wrong:
				let = let.Background(GRAY).Foreground(BLACK)
			case Partial:
				let = let.Background(YELLOW)
			case Correct:
				let = let.Background(GREEN)
			}
			word = lipgloss.JoinHorizontal(
				lipgloss.Left,
				word,
				let.Render(string(char.char)),
			)
		}
		row := rowStyle.Render(word)

		if result == "" {
			result = row
		} else {
			result = lipgloss.JoinVertical(lipgloss.Left, result, row)
		}
	}
	return result
}

func (m model) View() tea.View {
	// dELTE test
	var guesses string
	for _, word := range m.guesses {
		guesses += "\n"
		for _, char := range word {
			guesses += string(char.char)
		}
	}
	testString := fmt.Sprintf("%s\n", letterToString(m.answer)) + m.renderWords()

	// Init view
	v := tea.NewView(testString)
	v.AltScreen = true
	return v
}

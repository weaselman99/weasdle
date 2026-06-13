package main

import (
	"fmt"

	"charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// ---------------- STYLES	----------------
// Colours
var (
	WHITE  = lipgloss.Color("#dddddd")
	BLACK  = lipgloss.Black
	GRAY   = lipgloss.Color("#444444")
	RED    = lipgloss.Color("#ff5d62")
	GREEN  = lipgloss.Color("#98bb6c")
	YELLOW = lipgloss.Color("#ffa066")
)

var areaStyle = lipgloss.NewStyle().
	Width(60).
	Height(30).
	Border(lipgloss.RoundedBorder(), true).
	Align(lipgloss.Center).Margin(1, 4)

var textStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(WHITE)

var boxStyle = lipgloss.NewStyle().
	Inherit(textStyle).
	Margin(0, 1, 0).
	Padding(0, 2).
	Border(lipgloss.ThickBorder(), true)

var rowStyle = lipgloss.NewStyle().
	Width(56)

var lettersStyle = lipgloss.NewStyle().
	Margin(0, 0, 0).
	Padding(0, 1).
	Foreground(WHITE)

var keyboardStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder(), true)

var titleStyle = lipgloss.NewStyle().
	Bold(true).
	Italic(true).
	Border(lipgloss.NormalBorder()).
	BorderForeground(WHITE).
	Foreground(WHITE).
	Padding(0, 1)

var errorStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(RED).
	MarginLeft(2)

// ---------------- VIEWS	----------------

func renderWord(word []letter) string {
	var result string
	for _, char := range word {
		// Change BG based on the state
		let := boxStyle
		switch char.state {
		case Empty:
			let = let.Foreground(GRAY).BorderForeground(GRAY)
		case Wrong:
			let = let.Foreground(GRAY).BorderForeground(GRAY)
		case Partial:
			let = let.Foreground(YELLOW).BorderForeground(YELLOW)
		case Correct:
			let = let.Foreground(GREEN).BorderForeground(GREEN)
		}
		result = lipgloss.JoinHorizontal(
			lipgloss.Left,
			result,
			let.Render(string(char.char)),
		)
	}
	return result
}

func (m model) renderRows() string {
	var result string
	for i, word := range m.guesses {
		row := lipgloss.JoinHorizontal(
			lipgloss.Center,
			textStyle.Render(fmt.Sprintf("%d   ", i+1)),
			renderWord(word),
		)
		result = lipgloss.JoinVertical(lipgloss.Left, result, row)
	}
	return result
}

func (m model) renderKeys() string {
	var result string
	keys := [][]rune{
		{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'},
		{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'},
		{'z', 'x', 'c', 'v', 'b', 'n', 'm'},
	}

	for _, row := range keys {
		var rowString string
		for _, key := range row {
			let := m.letters[key]
			letStyle := lettersStyle
			switch let.state {
			case Wrong:
				letStyle = letStyle.Background(GRAY).Foreground(BLACK)
			case Partial:
				letStyle = letStyle.Background(YELLOW).Foreground(BLACK)
			case Correct:
				letStyle = letStyle.Background(GREEN).Foreground(BLACK)
			}
			rowString = lipgloss.JoinHorizontal(lipgloss.Center, rowString, letStyle.Render(string(let.char)))
		}
		if result == "" {
			result = rowString
		} else {
			result = lipgloss.JoinVertical(lipgloss.Center, result, rowString)
		}
	}
	return keyboardStyle.Render(result)
}

func (m model) renderTitle() string {
	var result string
	title := titleStyle.Render("Weasdle")
	errorString := errorStyle.Render(m.errorMsg)
	if m.gameState == Won {
		if m.wordPos == 0 {
			errorString = errorStyle.Foreground(GREEN).Render("You won in 1 try!")
		} else {
			errorString = errorStyle.Foreground(GREEN).Render(fmt.Sprintf("You won in %d tries!", m.wordPos+1))
		}
	} else if m.gameState == Lost {
		errorString = errorStyle.Render("You lost! The answer was " + fmt.Sprintf("\"%s\"", letterToString(m.answer)))
	}
	result = lipgloss.NewStyle().Width(52).Render(
		lipgloss.JoinHorizontal(lipgloss.Center, title, errorString),
	)

	return result
}

func (m model) View() tea.View {
	result := areaStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.renderTitle(),
			m.renderRows(),
			m.renderKeys(),
			// fmt.Sprintf("%s\n", letterToString(m.answer)), // Show answer
		))

	// Init view
	v := tea.NewView(result)
	v.AltScreen = true
	return v
}

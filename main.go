package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type GameState int

const (
	ModeGameStart GameState = iota
	ModeGamePlaying
	ModeGameOver
)

// model stores our application's state
type model struct {
	tea.Model
	gameState GameState
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Checks if a key is pressed
	case tea.KeyMsg:
		// What key was pressed?
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	}
	return m, nil
}

// The UI is a string
func (m model) View() string {
	// The header
	// The footer
	s := "Black Jack\n\n"

	s += "\nPress q to quit\n\n"

	// Send the UI for rendering
	return s
}

func initialModel() model {
	return model{
		gameState: ModeGameStart}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

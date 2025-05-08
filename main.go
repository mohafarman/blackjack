package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// model stores our application's state
type model struct {
	tea.Model
	blackjack blackJack
	width     int
	height    int
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("Black Jack"),
		tea.EnterAltScreen,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	// Checks if a key is pressed
	case tea.KeyMsg:
		// What key was pressed?
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}
	}

	return m, nil
}

// The UI is a string
func (m model) View() string {
	// The header
	s := "Black Jack\n\n"

	// The footer
	s += "\nPress q to quit\n\n"

	// Send the UI for rendering
	return s
}

func initialModel() model {
	return model{
		// TODO: Shuffle the deck
		blackjack: newGame(),
	}
}

func main() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatalf("Black Jack error: %v", err)
	}
}

package main

import (
	deck "blackjack/src"
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
	altscreen bool
	gameState GameState
	width     int
	height    int
	deck      deck.Deck
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
		gameState: ModeGameStart,
		// TODO: Allow user to specify amount of decks to play with
		deck: deck.NewDeck(1)}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

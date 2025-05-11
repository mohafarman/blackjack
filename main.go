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

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		// TODO: Handle Hit
		case "h":
			if m.blackjack.gameState == ModeGameStart {
				m.blackjack.playerHit()
			} else if m.blackjack.gameState == ModeGameOver {
				/* TODO: Continue playing with a new hand */
			}
			return m, nil

		// TODO: Handle Stand
		case "s":
			if m.blackjack.gameState == ModeGameStart {
				m.blackjack.gameState = ModeGameOver
				m.blackjack.determineWinner()
			}
			return m, nil

		case "r":
			// tea.WindowSize() allows for the RenderTitleRow to continue rendering properly
			return initialModel(), tea.Batch(tea.ClearScreen, tea.WindowSize())
		}

	}

	return m, nil
}

// The UI is a string
func (m model) View() string {
	return renderGameState(m.blackjack, m.width)
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

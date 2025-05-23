package main

import (
	"flag"
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
		tea.ClearScreen,
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

		case "h":
			if m.blackjack.gameState == ModeGameStart {
				m.blackjack.playerHit()
			} else if m.blackjack.gameState == ModeGameOver {
				newRound(&m.blackjack)
			}
			return m, tea.Batch(tea.ClearScreen, tea.WindowSize())

		case "s":
			if m.blackjack.gameState == ModeGameStart {
				m.blackjack.gameState = ModeGameOver
				m.blackjack.dealerPlay()
				m.blackjack.playerScore, _ = m.blackjack.calculateHandScore(m.blackjack.playerHand)
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
	return renderGameState(m, m.width)
}

func initialModel() model {
	log.Println("Number of decks:", *numDecks)
	log.Println("Hit on soft 17:", *h17)
	return model{
		blackjack: newGame(),
	}
}

var numDecks = flag.Int("decks", 4, "Number of decks to play with. 4-8 decks allowed.")
var h17 = flag.Bool("h17", true, "Dealer hits on soft 17. To set false, h17=false.")

func main() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	flag.Parse()

	if *numDecks < 4 || *numDecks > 8 {
		var Usage = func() {
			flag.PrintDefaults()
		}
		Usage()
		os.Exit(1)
	}

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatalf("Black Jack error: %v", err)
	}
}

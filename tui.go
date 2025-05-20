package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

// termenv for color styling

func renderCard(card Card) string {
	var p = termenv.ColorProfile()
	var suitSymbol string
	var suitColor termenv.Color

	switch card.Suit {
	case "Hearts":
		suitSymbol = "♥"
		suitColor = p.Color("#FF0000") // Red color
	case "Diamonds":
		suitSymbol = "♦"
		suitColor = p.Color("#FF0000") // Red color
	case "Clubs":
		suitSymbol = "♣"
		suitColor = p.Color("#FFFFFF") // Black color
	case "Spades":
		suitSymbol = "♠"
		suitColor = p.Color("#FFFFFF") // Black color
	}

	coloredSymbol := termenv.String(suitSymbol).Foreground(suitColor).String()
	return fmt.Sprintf("%s%s", card.Rank, coloredSymbol)
}

func renderHand(doc *strings.Builder, hand []Card) {
	for i := range hand {
		doc.WriteString(renderCard(hand[i]))
		doc.WriteString(" ")
	}
}

func renderScore(score int, soft bool) string {
	scoreText := fmt.Sprintf("%d", score)
	if soft {
		return fmt.Sprintf("\tSoft %s", scoreText)
	}
	return fmt.Sprintf("\t== %s", scoreText)
}

func renderGameState(m model, width int) string {
	doc := &strings.Builder{}
	// The header
	RenderTitleRow(width, doc, TitleRowProps{"Black Jack"})
	// At least one "\n" needs to be printed otherwise the next doc.WriteString is not displayed
	doc.WriteString("\n\n")

	/* Dealer hand */
	doc.WriteString("Dealer hand:\t")
	/* Hide one card when game is on */
	if m.blackjack.gameState == ModeGameStart {
		// Dealer score not be calculated right now. Only at the end of the game
		doc.WriteString("??")
		doc.WriteString(renderCard(m.blackjack.dealerHand[1]))
	}

	if m.blackjack.gameState == ModeGameOver {
		renderHand(doc, m.blackjack.dealerHand)
		doc.WriteString("\t")
		doc.WriteString(renderScore(m.blackjack.calculateHandScore(m.blackjack.dealerHand)))
	}
	doc.WriteString("\n\n")

	/* Player hand */
	doc.WriteString("Your hand:\t")
	renderHand(doc, m.blackjack.playerHand)
	doc.WriteString(renderScore(m.blackjack.calculateHandScore(m.blackjack.playerHand)))

	if m.blackjack.gameState == ModeGameOver {
		doc.WriteString("\n\n\n")
		doc.WriteString(gameOverMessage(m.blackjack))
		doc.WriteString("\n\n\nPress 'H' to play next hand\n\n")
	}

	// The footer
	if m.blackjack.gameState == ModeGameStart {
		doc.WriteString("\n\n\nPress 'H' to hit or 'S' to stand\n\n")
	}

	doc.WriteString("\nPress 'Q' to quit\n\n")

	// Send the UI for rendering
	return doc.String()
}

func gameOverMessage(bj blackJack) string {
	if bj.playerWins {
		return "Player Wins!"
	} else if bj.tie {
		return "Tie!"
	} else {
		return "Dealer Wins!"
	}
}

/***************************************************************************************************/
/* Code from https://github.com/doppelganger113/example-go-tui/blob/main/internal/tui/title_row.go */
/***************************************************************************************************/
type TitleRowProps struct {
	Title string
}

func RenderTitleRow(width int, doc *strings.Builder, props TitleRowProps) {
	var (
		highlight = lipgloss.AdaptiveColor{Light: "#347aeb", Dark: "#347aeb"}

		activeTabBorder = lipgloss.Border{
			Top:         "─",
			Bottom:      " ",
			Left:        "│",
			Right:       "│",
			TopLeft:     "╭",
			TopRight:    "╮",
			BottomLeft:  "┘",
			BottomRight: "└",
		}

		tabBorder = lipgloss.Border{
			Top:         "─",
			Bottom:      "─",
			Left:        "│",
			Right:       "│",
			TopLeft:     "╭",
			TopRight:    "╮",
			BottomLeft:  "┴",
			BottomRight: "┴",
		}

		tab = lipgloss.NewStyle().
			Border(tabBorder, true).
			BorderForeground(highlight).
			Padding(0, 1)

		activeTab = tab.Border(activeTabBorder, true)

		tabGap = tab.
			BorderTop(false).
			BorderLeft(false).
			BorderRight(false)
	)

	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		activeTab.Render(props.Title),
	)

	gap := tabGap.Render(strings.Repeat(" ", max(0, width-lipgloss.Width(row)-2)))
	row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)

	doc.WriteString(row)
}

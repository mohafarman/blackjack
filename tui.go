package main

import (
	"fmt"
	"log"
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

func renderGameState(bj blackJack, width int) string {
	doc := &strings.Builder{}
	// The header
	RenderTitleRow(width, doc, TitleRowProps{"Black Jack"})
	// At least one "\n" needs to be printed otherwise the next doc.WriteString is not displayed
	doc.WriteString("\n\n")

	if bj.gameState == ModeGameStart {
		/* TODO: Display the cards */
		/* Dealer hand */
		doc.WriteString("Dealer hand:\t")
		renderHand(doc, bj.dealerHand)

		doc.WriteString("\n\n")

		/* Player hand */
		doc.WriteString("Your hand:\t")
		renderHand(doc, bj.dealerHand)
		/* TODO: Render players hand score */
	}

	// The footer
	doc.WriteString("\n\nPress q to quit\n\n")

	log.Println()

	// Send the UI for rendering
	return doc.String()
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

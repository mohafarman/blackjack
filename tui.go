package main

import (
	"github.com/charmbracelet/lipgloss"
	"strings"
)

func renderGameState(gs GameState, width int) string {
	doc := &strings.Builder{}
	// The header
	RenderTitleRow(width, doc, TitleRowProps{"Black Jack\n\n"})

	if gs == ModeGameStart {
		/* TODO: Display the cards */
	}

	// The footer
	doc.WriteString("\nPress q to quit\n\n")

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

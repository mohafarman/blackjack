// logic.go handles the game logic
package main

import ()

type GameState int

const (
	ModeGameStart GameState = iota
	ModeGamePlaying
	ModeGameOver
)

type blackJack struct {
	gameState   GameState
	deck        Deck
	playerHand  []Card
	dealerHand  []Card
	playerScore int
	dealerScore int
}

func (bj *blackJack) dealInitCards() {
	var card Card
	for range 2 {
		/* Slicing slices to deal cards, deals one card at a time */
		card = bj.deck.Cards[0]
		bj.playerHand = append(bj.playerHand, card)
		bj.deck.Cards = bj.deck.Cards[1:]

		card = bj.deck.Cards[0]
		bj.dealerHand = append(bj.dealerHand, card)
		bj.deck.Cards = bj.deck.Cards[1:]
	}
}

func (bj blackJack) calculateHandScore(hand []Card) (int, bool) {
	var score, aces int
	var soft bool

	for _, card := range hand {
		switch {
		case card.Rank == "J" || card.Rank == "Q" || card.Rank == "K":
			score += 10
		case card.Rank == "A":
			aces++
			score += 11
		case card.Rank == "10":
			score += 10
		default:
			score += int(card.Rank[0] - '0') // Here using ASCII numbers to convert to int
		}
	}

	// Adjust for aces
	if score > 21 && aces > 0 {
		aces--
		score -= 10
		soft = false
	}

	if score <= 21 && aces > 0 {
		soft = true
	}

	return score, soft
}

func newGame() blackJack {
	bj := blackJack{
		gameState: ModeGameStart,
		// TODO: Allow user to specify amount of decks to play with
		deck: shuffleDeck(newDeck(4)),
	}

	bj.dealInitCards()
	return bj
}

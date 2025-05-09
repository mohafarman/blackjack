// logic.go handles the game logic
package main

type GameState int

const (
	ModeGameStart GameState = iota
	ModeGamePlaying
	ModeGameOver
)

type blackJack struct {
	gameState  GameState
	deck       Deck
	playerHand []Card
	dealerHand []Card
}

func (bj blackJack) dealInitCards() {
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

func newGame() blackJack {
	bj := blackJack{
		gameState: ModeGameStart,
		// TODO: Allow user to specify amount of decks to play with
		deck: shuffleDeck(newDeck(1)),
	}

	bj.dealInitCards()
	return bj
}

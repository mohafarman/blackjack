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

func newGame() blackJack {
	return blackJack{
		gameState: ModeGameStart,
		// TODO: Allow user to specify amount of decks to play with
		deck: shuffleDeck(newDeck(1))}
}

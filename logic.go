// logic.go handles the game logic
package main

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
	playerWins  bool
	tie         bool
}

func (bj *blackJack) dealInitCards() {
	var card Card
	bj.playerHand = []Card{}
	bj.dealerHand = []Card{}
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

func newRound(bj *blackJack) {
	bj.gameState = ModeGameStart
	bj.playerScore = 0
	bj.dealerScore = 0
	bj.dealInitCards()
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

func (bj *blackJack) playerHit() {
	var card Card
	card = bj.deck.Cards[0]
	bj.playerHand = append(bj.playerHand, card)
	bj.deck.Cards = bj.deck.Cards[1:]

	bj.playerScore, _ = bj.calculateHandScore(bj.playerHand)
	if bj.playerScore > 21 {
		bj.gameState = ModeGameOver
		bj.playerWins = false
	}
}

func (bj *blackJack) dealerHit() {
	var card Card
	card = bj.deck.Cards[0]
	bj.dealerHand = append(bj.dealerHand, card)
	bj.deck.Cards = bj.deck.Cards[1:]

	bj.playerScore, _ = bj.calculateHandScore(bj.dealerHand)
	if bj.dealerScore > 21 {
		bj.gameState = ModeGameOver
		bj.playerWins = true
	}
}

func (bj *blackJack) determineWinner() {
	bj.gameState = ModeGameOver

	bj.playerScore, _ = bj.calculateHandScore(bj.playerHand)
	bj.dealerScore, _ = bj.calculateHandScore(bj.dealerHand)

	if bj.playerScore > 21 {
		bj.playerWins = false
	} else if bj.dealerScore > 21 || bj.playerScore > bj.dealerScore {
		bj.playerWins = true
	} else if bj.dealerScore > bj.playerScore {
		bj.playerWins = false
	} else {
		bj.tie = true
		bj.playerWins = false
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

// Package provides a way to deal with a deck of cards
package deck

type Card struct {
	Suit string
	Rank string
}

type Deck struct {
	Cards []Card
}

func NewDeck(decks int) Deck {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	var deck []Card

	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}

	return Deck{Cards: deck}
}

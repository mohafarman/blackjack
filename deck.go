// Package provides a way to deal with a deck of cards
package main

import (
	"math/rand"
	"time"
)

type Card struct {
	Suit string
	Rank string
}

type Deck struct {
	Cards []Card
}

func newDeck(decks int) Deck {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	var deck []Card

	for range decks {
		for _, suit := range suits {
			for _, rank := range ranks {
				deck = append(deck, Card{Suit: suit, Rank: rank})
			}
		}
	}

	return Deck{Cards: deck}
}

func shuffleDeck(deck Deck) Deck {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(deck.Cards), func(i int, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
	return deck
}

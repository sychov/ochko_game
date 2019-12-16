package objects

import (
	"math/rand"
	"time"
)

// CardsNumber : number of cards in deck
const CardsNumber = 36

var suits = [4]string{"♠", "♥", "♣", "♦"}
var values = [9]string{"6", "7", "8", "9", "10", "В", "Д", "К", "Т"}

// Deck of players cards
type Deck struct {
	cards []Card
}

// Shuffle the Deck cards
func (deck *Deck) Shuffle() {
	deck.refill()
	nsec := time.Now().Nanosecond()
	rand.Seed(int64(nsec))
	rand.Shuffle(
		len(deck.cards),
		func(a, b int) {
			deck.cards[a], deck.cards[b] = deck.cards[b], deck.cards[a]
		},
	)
}

// Take card from the top of deck to hand
func (deck *Deck) Take(hand *Hand) {
	hand.Add(deck.cards[1])
	deck.cards = deck.cards[1:]
}

// refill deck with cards. Private.
func (deck *Deck) refill() {
	deck.cards = nil
	for q := 0; q < CardsNumber; q++ {
		deck.cards = append(
			deck.cards,
			NewCard(suits[q%4], values[q/4]),
		)
	}
}

// NewDeck : Deck's factory
func NewDeck() Deck {
	var deck Deck
	deck.Shuffle()
	return deck
}

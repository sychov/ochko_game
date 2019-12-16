package objects

import "fmt"

// Hand cards of player
type Hand struct {
	cards []Card
}

// Calculate : return points of all cards in the hand
func (hand *Hand) Calculate() int {
	var points int
	for _, card := range hand.cards {
		points += card.points
	}
	return points
}

// Add card to players hand
func (hand *Hand) Add(card Card) {
	hand.cards = append(hand.cards, card)
}

// GetRepresentation : get string representation of the hand
func (hand *Hand) GetRepresentation() string {
	var representation string
	for index, card := range hand.cards {
		if index > 0 {
			representation += ", "
		}
		representation += card.title
	}
	suffix := ""
	if hand.Calculate() > 21 {
		suffix = " - перебор"
	}
	return fmt.Sprintf("%s (%d очков%s).\n", representation, hand.Calculate(), suffix)
}

// GetInitialHand : create a hand and take two cards from the deck to it
func GetInitialHand(deck *Deck) Hand {
	hand := Hand{}
	deck.Take(&hand)
	deck.Take(&hand)
	return hand
}

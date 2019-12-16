package objects

var points = map[string]int{
	"6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "В": 2, "Д": 3, "К": 4, "Т": 11,
}

// Card of players deck's
type Card struct {
	title  string
	points int
}

// NewCard : Card's factory
func NewCard(suit, value string) Card {
	return Card{
		title:  value + suit,
		points: points[value],
	}
}

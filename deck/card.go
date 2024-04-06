package deck

var cardVals = [...]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
var cardTypes = [...]string{"Hearts", "Diamonds", "Spades", "Clubs"}

type Card struct {
	Value string
	Type  string
}

func New() []Card {
	var deck []Card
	for _, x := range cardTypes {
		for i := 0; i < len(cardVals); i++ {
			newCard := Card{
				Value: cardVals[i],
				Type:  x,
			}
			deck = append(deck, newCard)
		}
	}
	return deck
}

func sortDeck() {}

func compare() {}

func shuffle() {}

func addJokers() {}

func filterCards() {}

func multiDeck() {}

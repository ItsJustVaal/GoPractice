package deck

var cardSymbols = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
var cardVals = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
var cardTypes = []string{"Hearts", "Diamonds", "Spades", "Clubs"}

type Card struct {
	Symbol string
	Value  int
	Type   string
}

type Deck struct {
	Cards []Card
}

func New() Deck {
	var deck Deck
	for _, x := range cardTypes {
		for i := 0; i < len(cardSymbols); i++ {
			newCard := Card{
				Symbol: cardSymbols[i],
				Value:  cardVals[i],
				Type:   x,
			}
			deck.Cards = append(deck.Cards, newCard)
		}
	}
	return deck
}

// func (d *Deck) sortDeck() {}

func Compare(firstCard, secondCard Card) Card {
	if firstCard.Value == secondCard.Value {
		return Card{
			Value: firstCard.Value,
		}
	}
	if firstCard.Value > secondCard.Value {
		return firstCard
	}
	if secondCard.Value > firstCard.Value {
		return secondCard
	}
	return Card{}
}

// func (d *Deck) shuffle() {}

func (d *Deck) AddJokers() {
	for _, x := range cardTypes {
		d.Cards = append(d.Cards, Card{
			Symbol: "Joker",
			Value:  15,
			Type:   x,
		})
	}
}

// func (d *Deck) filterCards() {}

// func (d *Deck) multiDeck() {}

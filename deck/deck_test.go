package deck

import "testing"

func TestCompare(t *testing.T) {
	cardOne := Card{
		Symbol: "Q",
		Value:  12,
		Type:   "Hearts",
	}
	cardTwo := Card{
		Symbol: "J",
		Value:  11,
		Type:   "Spades",
	}
	cardThree := Card{
		Symbol: "Q",
		Value:  12,
		Type:   "Jacks",
	}
	cardFour := Card{
		Symbol: "A",
		Value:  14,
		Type:   "Diamonds",
	}

	cardOneWins := Compare(cardOne, cardTwo)
	if cardOneWins != cardOne {
		t.Errorf("Expected cardOne, got: %v", cardOneWins)
	}
	cardTwoWins := Compare(cardThree, cardFour)
	if cardTwoWins != cardFour {
		t.Errorf("Expected cardFour, got: %v", cardTwoWins)
	}
	cardThreeWins := Compare(cardFour, cardThree)
	if cardThreeWins != cardFour {
		t.Errorf("Expected cardFour, got: %v", cardThreeWins)
	}
	tie := Compare(cardOne, cardThree)
	if tie.Value != cardOne.Value {
		t.Errorf("Expected Tie, got: %v", tie)
	}
}

func TestNew(t *testing.T) {
	deck := New()
	if len(deck.Cards) != 56 {
		t.Errorf("Not a complete deck, only got %d cards instead of 56", len(deck.Cards))
	}
}

func TestAddJokers(t *testing.T) {
	deck := New()
	deck.AddJokers()
	var numJokers int
	for _, card := range deck.Cards {
		if card.Symbol == "Joker" {
			numJokers++
		}
	}
	if numJokers != 4 {
		t.Errorf("Expected 4 Jokers to be added, got: %d", numJokers)
	}
}

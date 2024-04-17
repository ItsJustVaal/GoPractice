package deck

import (
	"testing"
)

func TestNew(t *testing.T) {
	deck := New()
	if len(deck.Cards) != 13*4 {
		t.Errorf("Wrong number of cards in deck %d", len(deck.Cards))
	}
}

func TestDefaultSort(t *testing.T) {
	d := New(DefaultSort)
	temp := Card{Rank: Ace, Suit: Spade}
	if d.Cards[0] != temp {
		t.Errorf("Expected Ace of Spades. Recieved: %s", d.Cards[0].String())
	}
}

// TestSort is to test custom sort functions as an opt
// The current test uses Less and expects the same outcome
// as TestDefaultSort
func TestSort(t *testing.T) {
	d := New(Sort(Less))
	temp := Card{Rank: Ace, Suit: Spade}
	if d.Cards[0] != temp {
		t.Errorf("Expected Ace of Spades. Recieved: %s", d.Cards[0].String())
	}
}

func TestAddJokers(t *testing.T) {
	d := New(AddJokers(3))
	count := 0
	for _, c := range d.Cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("Expected 3 Jokers, Recieved: $d", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == 3
	}
	d := New(Filter(filter))
	for _, c := range d.Cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected twos and threes to be filtered out.")
		}
	}

}

func TestAddDeck(t *testing.T) {
	toAdd := 4
	d := New(AddDeck(toAdd))
	if len(d.Cards) != (toAdd+1)*(13*4) {
		t.Errorf("Expected 104 cards, got %d", len(d.Cards))
	}
}

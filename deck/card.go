//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // Special Case
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit Suit
	Rank Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New(opts ...func([]Card) []Card) []Card {
	var d []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			d = append(d, Card{Suit: suit, Rank: rank})
		}
	}
	for _, opt := range opts {
		d = opt(d)
	}
	return d
}

func DefaultSort(d []Card) []Card {
	sort.Slice(d, Less(d))
	return d
}

func Sort(less func(cards []Card) func(i, j int) bool) func(d []Card) []Card {
	return func(d []Card) []Card {
		sort.Slice(d, Less(d))
		return d
	}
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

// Shuffle is a simple shuffle func that uses
// a random perm to make a new slice of shuffled cards
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i, j := range r.Perm(len(cards)) {
		ret[i] = cards[j]
	}
	return ret
}

func AddJokers(toAdd int) func(d []Card) []Card {
	return func(d []Card) []Card {
		for i := 0; i < toAdd; i++ {
			d = append(d, Card{
				Rank: Rank(i),
				Suit: Joker,
			})
		}
		return d
	}
}

func Filter(f func(card Card) bool) func(d []Card) []Card {
	return func(d []Card) []Card {
		var ret []Card
		for _, c := range d {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		d = ret
		return d
	}
}

// AddDeck duplicates all current cards in the deck
// run this after your first deck is sorted, filtered
// with how many jokers you want before using this
func AddDeck(toAdd int) func([]Card) []Card {
	return func(d []Card) []Card {
		var ret []Card
		for i := 0; i < toAdd+1; i++ {
			ret = append(ret, d...)
		}
		d = ret
		return d
	}
}

package main

import (
	"fmt"
	"strings"

	"github.com/ItsJustVaal/GoPractice/deck"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealersHand() string {
	return h[0].String() + ", **HIDDEN**"
}

func (h Hand) Score() int {
	minScore := h.minScore()
	if minScore > 11 {
		return minScore
	}

	for _, c := range h {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}

	return minScore
}

func (h Hand) minScore() int {
	score := 0
	for _, card := range h {
		score += min(int(card.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	ret.Deck = deck.New(deck.AddDeck(3))
	return ret
}

func Deal(gs GameState) GameState {
	ret := clone(gs)
	// Hand is a []deck.Card
	ret.Player = make(Hand, 0, 5)
	ret.Dealer = make(Hand, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, ret.Deck = draw(ret.Deck)
		ret.Player = append(ret.Player, card)
		card, ret.Deck = draw(ret.Deck)
		ret.Dealer = append(ret.Dealer, card)
	}
	ret.State = StatePlayerTurn
	return ret
}

func Stand(gs GameState) GameState {
	ret := clone(gs)
	ret.State++

	return ret
}

func EndHand(gs GameState) GameState {
	ret := clone(gs)
	pScore, dScore := ret.Player.Score(), ret.Dealer.Score()
	fmt.Println()
	switch {
	case pScore > 21:
		fmt.Println("You Busted")
	case dScore > 21:
		fmt.Println("Dealer Busted")
	case pScore > dScore:
		fmt.Println("You Win")
	case dScore > pScore:
		fmt.Println("Dealer Wins")
	case dScore == pScore:
		fmt.Println("Draw")
	}
	fmt.Println("==== Final Hands ====")
	fmt.Printf("Player Hand: %s\nPlayer Score: %d\n", ret.Player, pScore)
	fmt.Printf("Dealer Hand: %s\nDealer Score: %d\n", ret.Dealer, dScore)
	fmt.Println()
	ret.Player = nil
	ret.Dealer = nil
	return ret
}

func Hit(gs GameState) GameState {
	ret := clone(gs)
	hand := ret.CurrentPlayer()
	var card deck.Card
	card, ret.Deck = draw(ret.Deck)
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(ret)
	}
	return ret
}

func New() Game {

}

func main() {
	var gs GameState
	gs = Shuffle(gs)

	// Gameplay loop starts here
	for i := 0; i < 10; i++ {
		gs = Deal(gs)
		// Player Turn Start
		var input string
		for gs.State == StatePlayerTurn {
			fmt.Println("Player Hand: ", gs.Player)
			fmt.Println("Dealer Hand: ", gs.Dealer.DealersHand())
			fmt.Println("Would you like to (h)it or (s)tand?")
			fmt.Scanf("%s\n", &input)
			switch input {
			case "h":
				gs = Hit(gs)
			case "s":
				gs = Stand(gs)
			case "q":
				fmt.Println()
				fmt.Println("Successfully Quit")
				return
			default:
				fmt.Println("Invalid Option")
			}
		}

		// Dealers Turn
		for gs.State == StateDealerTurn {
			if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.minScore() != 17) {
				gs = Hit(gs)
			} else {
				gs = Stand(gs)
			}
		}

		gs = EndHand(gs)
	}
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

type State uint8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type GameState struct {
	Deck   []deck.Card
	State  State
	Player Hand
	Dealer Hand
}

func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("It isnt current any players turn")
	}
}

func clone(gs GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make([]deck.Card, len(gs.Player)),
		Dealer: make([]deck.Card, len(gs.Dealer)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)
	return ret
}

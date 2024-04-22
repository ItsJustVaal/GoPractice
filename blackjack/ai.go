package blackjack

import (
	"fmt"

	"github.com/ItsJustVaal/GoPractice/deck"
)

type AI interface {
	Bet() int
	Play(hand []deck.Card, dealer deck.Card) Move
	Results(hand [][]deck.Card, dealer []deck.Card)
}

type HumanAI struct{}

func (ai *HumanAI) Bet() int {
	return 1
}

func (ai *HumanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println("Player Hand: ", hand)
		fmt.Println("Dealer Hand: ", dealer)
		fmt.Println("Would you like to (h)it or (s)tand?")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return Hit
		case "s":
			return Stand
		default:
			fmt.Println("Invalid Option")
		}
	}
}

func (ai *HumanAI) Results(hand [][]deck.Card, dealer []deck.Card) {

}

type Move func(GameState) GameState

type GameState struct{}

func Hit(gs GameState) GameState {
	return gs
}

func Stand(gs GameState) GameState {
	return gs
}

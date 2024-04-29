package blackjack

import (
	"fmt"

	"github.com/ItsJustVaal/GoPractice/deck"
)

type AI interface {
	Bet(shuffled bool) int
	Play(hand []deck.Card, dealer deck.Card) Move
	Results(hand [][]deck.Card, dealer []deck.Card)
}

type dealerAI struct{}

func (ai *dealerAI) Bet(shuffled bool) int {
	if shuffled {
		fmt.Println("Deck was just shuffled")
	}
	return 1
}

func (ai dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	dScore := Score(hand...)
	if dScore <= 16 || (dScore == 17 && Soft(hand...)) {
		return MoveHit
	}
	return MoveStand
}

func (ai dealerAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	fmt.Println("==== Final Hands ====")
	fmt.Printf("Player Hand: %s\n", hand)
	fmt.Printf("Dealer Hand: %s\n", dealer)
	fmt.Println()
}

func HumanAI() AI {
	return humanAI{}
}

type humanAI struct{}

func (ai humanAI) Bet(shuffled bool) int {
	if shuffled {
		fmt.Println("Deck was just shuffled")
	}
	fmt.Println("What would you like to bet?")
	var bet int
	fmt.Scanf("%d\n", &bet)
	return bet
}

func (ai humanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println()
		fmt.Println("Player Hand: ", hand)
		fmt.Println("Dealer Hand: ", dealer)
		fmt.Println("Would you like to (h)it, (s)tand, s(p)lit or (d)ouble?")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		case "d":
			return MoveDouble
		case "p":
			return MoveSplit
		default:
			fmt.Println("Invalid Option")
		}
	}
}

func (ai humanAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	fmt.Println("==== Final Hands ====")
	fmt.Printf("Player Hand: %s\n", hand)
	fmt.Printf("Dealer Hand: %s\n", dealer)
	fmt.Println()
}

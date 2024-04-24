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

type dealerAI struct{}

func (ai *dealerAI) Bet() int {
	return 1
}

func (ai dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	dScore := score(hand...)
	if dScore <= 16 || (dScore == 17 && Soft(hand...)) {
		return MoveHit
	}
	return MoveStand
}

func (ai dealerAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	fmt.Println("==== Final Hands ====")
	fmt.Printf("Player Hand: %s\n", hand)
	fmt.Printf("Dealer Hand: %s\n", dealer)
}

func HumanAI() AI {
	return humanAI{}
}

type humanAI struct{}

func (ai humanAI) Bet() int {
	return 1
}

func (ai humanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println()
		fmt.Println("Player Hand: ", hand)
		fmt.Println("Dealer Hand: ", dealer)
		fmt.Println("Would you like to (h)it or (s)tand?")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		default:
			fmt.Println("Invalid Option")
		}
	}
}

func (ai humanAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	fmt.Println("==== Final Hands ====")
	fmt.Printf("Player Hand: %s\n", hand)
	fmt.Printf("Dealer Hand: %s\n", dealer)
}

package main

import (
	"fmt"

	"github.com/ItsJustVaal/GoPractice/blackjack"
	"github.com/ItsJustVaal/GoPractice/deck"
)

type basicAI struct {
	score int
	seen  int
	decks int
}

// Bet is not great because the AI isnt great
func (ai *basicAI) Bet(shuffled bool) int {
	if shuffled {
		ai.score = 0
		ai.seen = 0
	}
	// Not 100% sure why this needed to be a float but
	// it would hit a divide by 0 error after around 20 hands
	trueScore := float64(ai.score) / float64((ai.decks*52-ai.seen)/52)
	switch {
	case trueScore > 14:
		return 10000
	case trueScore > 8:
		return 500
	default:
		return 100
	}
}

// Play has an EXTREMELY simplified version of blackjack AI
func (ai *basicAI) Play(hand []deck.Card, dealer deck.Card) blackjack.Move {
	score := blackjack.Score(hand...)
	if len(hand) == 2 {
		if hand[0] == hand[1] {
			cardScore := blackjack.Score(hand[0])
			if cardScore >= 8 && cardScore != 10 {
				return blackjack.MoveSplit
			}
		}
		if (score == 10 || score == 11) && !blackjack.Soft(hand...) {
			return blackjack.MoveDouble
		}
	}
	dScore := blackjack.Score(dealer)
	if dScore >= 5 && dScore <= 6 {
		return blackjack.MoveStand
	}
	if score < 13 {
		return blackjack.MoveHit
	}
	return blackjack.MoveStand
}

func (ai *basicAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	for _, card := range dealer {
		ai.count(card)
	}
	for _, han := range hand {
		for _, cardz := range han {
			ai.count(cardz)
		}
	}
}

func (ai *basicAI) count(card deck.Card) {
	score := blackjack.Score(card)
	switch {
	case score >= 10:
		ai.score--
	case score <= 6:
		ai.score++
	}
	ai.seen++
}

func main() {
	opts := blackjack.Options{
		Decks:           4,
		Hands:           99999,
		BlackjackPayout: 1.5,
	}
	game := blackjack.New(opts)
	winnings := game.Play(&basicAI{
		decks: 4,
	})
	fmt.Printf("Winnings: %d\n", winnings)
}

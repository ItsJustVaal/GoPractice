package blackjack

import (
	"fmt"

	"github.com/ItsJustVaal/GoPractice/deck"
)

const (
	statePlayerTurn state = iota
	stateDealerTurn
	stateHandOver
)

type state uint8

type Game struct {
	deck     []deck.Card
	state    state
	player   []deck.Card
	dealer   []deck.Card
	dealerAI dealerAI
	balance  int
}

func New() Game {
	return Game{
		state:    statePlayerTurn,
		dealerAI: dealerAI{},
		balance:  0,
	}
}

func (g *Game) currentPlayer() *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player
	case stateDealerTurn:
		return &g.dealer
	default:
		panic("It isnt current any players turn")
	}
}

func deal(g *Game) {
	// Hand is a []deck.Card
	g.player = make([]deck.Card, 0, 5)
	g.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.deck = draw(g.deck)
		g.player = append(g.player, card)
		card, g.deck = draw(g.deck)
		g.dealer = append(g.dealer, card)
	}
	g.state = statePlayerTurn
}

func (g *Game) Play(ai AI) int {
	g.deck = deck.New(deck.AddDeck(3))

	for i := 0; i < 10; i++ {
		deal(g)
		for g.state == statePlayerTurn {
			hand := make([]deck.Card, len(g.player))
			copy(hand, g.player)
			move := ai.Play(g.player, g.dealer[0])
			move(g)
		}

		// Dealers Turn
		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.dealer))
			copy(hand, g.dealer)
			move := g.dealerAI.Play(hand, g.dealer[0])
			move(g)
		}

		endHand(g, ai)
	}
	return g.balance
}

type Move func(*Game)

func MoveHit(g *Game) {
	hand := g.currentPlayer()
	var card deck.Card
	card, g.deck = draw(g.deck)
	*hand = append(*hand, card)
	if score(*hand...) > 21 {
		MoveStand(g)
	}
}

func MoveStand(g *Game) {
	g.state++
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func endHand(g *Game, ai AI) {
	pScore, dScore := score(g.player...), score(g.dealer...)
	switch {
	case pScore > 21:
		fmt.Println("You Busted")
		g.balance--
	case dScore > 21:
		fmt.Println("Dealer Busted")
		g.balance++
	case pScore > dScore:
		fmt.Println("You Win")
		g.balance++
	case dScore > pScore:
		fmt.Println("Dealer Wins")
		g.balance--
	case dScore == pScore:
		fmt.Println("Draw")
	}
	fmt.Println()
	ai.Results([][]deck.Card{g.player}, g.dealer)
	g.player = nil
	g.dealer = nil
}

func score(hand ...deck.Card) int {
	minScore := minScore(hand...)
	if minScore > 11 {
		return minScore
	}

	for _, c := range hand {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}

	return minScore
}

func Soft(hand ...deck.Card) bool {
	minScore := minScore(hand...)
	score := score(hand...)
	return minScore != score
}

func minScore(hand ...deck.Card) int {
	score := 0
	for _, card := range hand {
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

package blackjack

import (
	"errors"

	"github.com/ItsJustVaal/GoPractice/deck"
)

const (
	statePlayerTurn state = iota
	stateDealerTurn
	stateHandOver
)

type state uint8

type Game struct {

	// Unexported opts
	nDecks          int
	nHands          int
	blackjackPayout float64

	state state
	deck  []deck.Card

	player    []hand
	handIdx   int
	playerBet int
	balance   int

	dealer   []deck.Card
	dealerAI dealerAI
}

type Options struct {
	Decks           int
	Hands           int
	BlackjackPayout float64
}

func New(opts Options) Game {
	g := Game{
		state:    statePlayerTurn,
		dealerAI: dealerAI{},
		balance:  0,
	}
	if opts.Decks == 0 {
		opts.Decks = 3
	}
	if opts.Hands == 0 {
		opts.Hands = 10
	}
	if opts.BlackjackPayout == 0 {
		opts.BlackjackPayout = 1.5
	}
	g.nDecks = opts.Decks
	g.nHands = opts.Hands
	g.blackjackPayout = opts.BlackjackPayout

	return g
}

func (g *Game) currentPlayer() *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player[g.handIdx].cards
	case stateDealerTurn:
		return &g.dealer
	default:
		panic("It isnt current any players turn")
	}
}

type hand struct {
	cards []deck.Card
	bet   int
}

func bet(g *Game, ai AI, shuffled bool) {
	bet := ai.Bet(shuffled)
	if bet < 100 {
		panic("bet must be at least 100")
	}
	g.playerBet = bet
}

func deal(g *Game) {
	// Hand is a []deck.Card
	playerHand := make([]deck.Card, 0, 5)
	g.handIdx = 0
	g.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.deck = draw(g.deck)
		playerHand = append(playerHand, card)
		card, g.deck = draw(g.deck)
		g.dealer = append(g.dealer, card)
	}
	g.player = []hand{
		{
			cards: playerHand,
			bet:   g.playerBet,
		},
	}
	g.state = statePlayerTurn
}

func (g *Game) Play(ai AI) int {
	g.deck = nil
	min := 52 * g.nDecks / 3
	for i := 0; i < g.nHands; i++ {
		shuffled := false
		for len(g.deck) < min {
			g.deck = deck.New(deck.AddDeck(g.nDecks), deck.Shuffle)
			shuffled = true
		}
		bet(g, ai, shuffled)
		deal(g)
		if Blackjack(g.dealer...) {
			endHand(g, ai)
			continue
		}
		for g.state == statePlayerTurn {
			hand := make([]deck.Card, len(*g.currentPlayer()))
			copy(hand, *g.currentPlayer())
			move := ai.Play(hand, g.dealer[0])
			err := move(g)
			switch err {
			case errorBust:
				MoveStand(g)
			case nil:
			default:
				panic(err)
			}
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

var (
	errorBust = errors.New("hand score exceeded 21")
)

type Move func(*Game) error

func MoveHit(g *Game) error {
	hand := g.currentPlayer()
	var card deck.Card
	card, g.deck = draw(g.deck)
	*hand = append(*hand, card)
	if Score(*hand...) > 21 {
		return errorBust
	}
	return nil
}

func MoveSplit(g *Game) error {
	cards := g.currentPlayer()
	if len(*cards) != 2 {
		return errors.New("you can only split with two cards in your hand")
	}
	if (*cards)[0].Rank != (*cards)[1].Rank {
		return errors.New("both cards must have the same rank to split")
	}
	g.player = append(g.player, hand{
		cards: []deck.Card{(*cards)[1]},
		bet:   g.player[g.handIdx].bet,
	})
	g.player[g.handIdx].cards = (*cards)[:1]
	return nil
}

func MoveDouble(g *Game) error {
	if len(*g.currentPlayer()) != 2 {
		return errors.New("can only double on a hand with 2 cards")
	}
	g.playerBet *= 2
	MoveHit(g)
	MoveStand(g)
	return nil
}

func MoveStand(g *Game) error {
	if g.state == stateDealerTurn {
		g.state++
		return nil
	}
	if g.state == statePlayerTurn {
		g.handIdx++
		if g.handIdx >= len(g.player) {
			g.state++
		}
		return nil
	}
	return errors.New("invalid state")
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func endHand(g *Game, ai AI) {
	dScore := Score(g.dealer...)
	dBJ := Blackjack(g.dealer...)
	allHands := make([][]deck.Card, len(g.player))
	for hidx, hand := range g.player {
		cards := hand.cards
		allHands[hidx] = cards
		pScore, pBJ := Score(cards...), Blackjack(cards...)
		winnings := hand.bet
		switch {
		case pBJ && dBJ:
			winnings = 0
		case dBJ:
			winnings = -winnings
		case pBJ:
			winnings = int(float64(winnings) * g.blackjackPayout)
		case pScore > 21:
			winnings = -winnings
		case dScore > 21:
		case pScore > dScore:
		case dScore > pScore:
			winnings = -winnings
		case dScore == pScore:
			winnings = 0
		}
		g.balance += winnings
	}
	ai.Results(allHands, g.dealer)
	g.player = nil
	g.dealer = nil
}

func Score(hand ...deck.Card) int {
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
	score := Score(hand...)
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

func Blackjack(hand ...deck.Card) bool {
	return len(hand) == 2 && Score(hand...) == 21
}

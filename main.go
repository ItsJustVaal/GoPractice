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

func main() {
	var card deck.Card
	deck := deck.New(deck.AddDeck(3))
	var player, dealer Hand
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, deck.Cards = Draw(deck.Cards)
			*hand = append(*hand, card)
		}
	}

	// Gameplay loop starts here
	var input string
	for input != "s" {
		fmt.Println("Player Hand: ", player)
		fmt.Println("Dealer Hand: ", dealer.DealersHand())
		fmt.Println("Would you like to (h)it or (s)tand?")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, deck.Cards = Draw(deck.Cards)
			player = append(player, card)
		}
	}
	fmt.Println("Final Hands")
	fmt.Println("Player Hand: ", player)
	fmt.Println("Dealer Hand: ", dealer)
}

func Draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

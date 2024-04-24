package main

import (
	"fmt"

	"github.com/ItsJustVaal/GoPractice/blackjack"
)

func main() {
	game := blackjack.New()
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}

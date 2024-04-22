package blackjack

import "github.com/ItsJustVaal/GoPractice/deck"

type AI interface{}

type HumanAI struct{}

func (ai *HumanAI) Play(hand []deck.Card) {

}

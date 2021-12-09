package dao

import (
	"github.com/google/uuid"
	"strings"
)

// list of suits
var Suits = [...]string{
	"SPADES",
	"DIAMONDS",
	"CLUBS",
	"HEARTS",
}

// list of card values
var Values = [...]string{
	"ACE",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"10",
	"JACK",
	"QUEEN",
	"KING",
}

type (
	Card struct {
		Value string `json:"value"`
		Suit  string `json:"suit"`
		Code  string `json:"code"`
	}

	ListCard struct {
		Cards []Card `json:"cards"`
	}
)

func (c *ListCard) FilterCards(codes []string) {
	var filteredCards []Card

	for _, card := range c.Cards {
		for _, code := range codes {
			if card.Code == strings.ToUpper(code) {
				filteredCards = append(filteredCards, card)
			}
		}
	}

	c.Cards = filteredCards
}

func (c *ListCard) ToCardDecks(deckID uuid.UUID) []CardDeck {
	var (
		seq       uint
		cardDecks []CardDeck
	)

	for _, card := range c.Cards {
		seq++
		cardDeck := CardDeck{
			DeckID:   deckID,
			Code:     card.Code,
			Sequence: seq,
			Value:    card.Value,
			Suit:     card.Suit,
			IsDrawn:  false,
		}
		cardDecks = append(cardDecks, cardDeck)
	}

	return cardDecks
}

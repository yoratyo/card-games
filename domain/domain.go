package domain

import "github.com/yoratyo/card-games/domain/deckOfCards"

type Domain struct {
	DeckOfCards deckOfCards.Service
}

func New(
	deckOfCards deckOfCards.Service,
) Domain {
	return Domain{
		DeckOfCards: deckOfCards,
	}
}

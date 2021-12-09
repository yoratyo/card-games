package handler

import "github.com/yoratyo/card-games/handler/deck"

type Handler struct {
	Deck deck.Handler
}

func New(deck deck.Handler) (Handler) {
	return Handler{
		Deck: deck,
	}
}

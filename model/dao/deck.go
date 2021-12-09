package dao

import (
	"github.com/google/uuid"
	"math/rand"
)

type (
	Deck struct {
		DeckID     uuid.UUID  `gorm:"primary_key"`
		IsShuffled bool       `gorm:"not_null"`
		Remaining  uint       `gorm:"not_null"`
		IsUsed     bool       `gorm:"not_null"`
		Cards      []CardDeck `gorm:"foreignKey:DeckID;references:DeckID"`
	}
)

func (Deck) TableName() string {
	return "decks"
}

func (Deck) ModelName() string {
	return "Deck"
}

func (d *Deck) Shuffle(cards ListCard) {
	rand.Shuffle(len(cards.Cards), func(i, j int) {
		cards.Cards[i], cards.Cards[j] = cards.Cards[j], cards.Cards[i]
	})

	d.Cards = cards.ToCardDecks(d.DeckID)
	d.IsShuffled = true
}

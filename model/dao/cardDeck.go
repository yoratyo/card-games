package dao

import "github.com/google/uuid"

type (
	CardDeck struct {
		DeckID   uuid.UUID `gorm:"primaryKey"`
		Code     string    `gorm:"primaryKey"`
		Sequence uint      `gorm:"not_null"`
		Value    string    `gorm:"not_null"`
		Suit     string    `gorm:"not_null"`
		IsDrawn  bool      `gorm:"not_null"`
	}
)

func (CardDeck) TableName() string {
	return "card_decks"
}

func (CardDeck) ModelName() string {
	return "CardDeck"
}

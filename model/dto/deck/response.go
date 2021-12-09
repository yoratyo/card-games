package deck

import (
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/model/dao"
)

type (
	CreateNewDeckResponseDTO struct {
		DeckID    uuid.UUID `json:"deck_id"`
		Shuffled  bool      `json:"shuffled"`
		Remaining uint      `json:"remaining"`
	}

	OpenDeckResponseDTO struct {
		DeckID    uuid.UUID `json:"deck_id"`
		Shuffled  bool      `json:"shuffled"`
		Remaining uint      `json:"remaining"`
		dao.ListCard
	}

	DrawCardsFromDeckResponseDTO struct {
		dao.ListCard
	}
)

package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/model/dao"
	"github.com/yoratyo/card-games/model/dto/deck"
)

func (s *service) GetDetailByID(ctx *gin.Context, ID uuid.UUID) (deck.OpenDeckResponseDTO, error) {
	deckDetail, err := s.repository.GetDeckDetailByID(ctx, ID)
	if err != nil {
		return deck.OpenDeckResponseDTO{}, err
	}

	cardDeck, err := s.repository.GetCardDeckByDeckID(ctx, ID)
	if err != nil {
		return deck.OpenDeckResponseDTO{}, err
	}

	var responseCardDetail []dao.Card

	for _, c := range cardDeck {
		card := dao.Card{
			Value: c.Value,
			Suit:  c.Suit,
			Code:  c.Code,
		}
		responseCardDetail = append(responseCardDetail, card)
	}

	return deck.OpenDeckResponseDTO{
		DeckID:    deckDetail.DeckID,
		Shuffled:  deckDetail.IsShuffled,
		Remaining: deckDetail.Remaining,
		ListCard:  dao.ListCard{Cards: responseCardDetail},
	}, nil
}

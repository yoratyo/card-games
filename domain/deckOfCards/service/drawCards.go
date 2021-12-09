package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/model/dao"
	"github.com/yoratyo/card-games/model/dto/deck"
	"gorm.io/gorm"
)

func (s *service) DrawCards(
	ctx *gin.Context,
	ID uuid.UUID,
	req deck.DrawCardsFromDeckRequestDTO,
) (deck.DrawCardsFromDeckResponseDTO, error) {
	count := int(req.Count)
	if count < 1 {
		count = 1
	}

	updateCardDeck := map[string]interface{}{
		"is_drawn": true,
	}

	cardDecks, err := s.repository.PatchDrawCardDeck(ctx, ID, count, updateCardDeck)
	if err != nil {
		return deck.DrawCardsFromDeckResponseDTO{}, err
	}

	if len(cardDecks) < count {
		return deck.DrawCardsFromDeckResponseDTO{}, errors.New("count param is more than remaining deck")
	}

	updateDeck := map[string]interface{}{
		"remaining": gorm.Expr("remaining - ?", count),
		"is_used":   true,
	}

	rowDeck, err := s.repository.PatchDeckByID(ctx, ID, updateDeck)
	if err != nil {
		return deck.DrawCardsFromDeckResponseDTO{}, err
	}

	if rowDeck != 1 {
		return deck.DrawCardsFromDeckResponseDTO{}, errors.New("failed to update remaining deck")
	}

	var responseDrawnCard []dao.Card

	for _, c := range cardDecks {
		card := dao.Card{
			Value: c.Value,
			Suit:  c.Suit,
			Code:  c.Code,
		}
		responseDrawnCard = append(responseDrawnCard, card)
	}

	return deck.DrawCardsFromDeckResponseDTO{
		ListCard: dao.ListCard{Cards: responseDrawnCard},
	}, nil
}

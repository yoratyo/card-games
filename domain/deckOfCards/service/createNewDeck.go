package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/card-games/model/dao"
	"github.com/yoratyo/card-games/model/dto/deck"
	"strings"
)

func (s *service) CreateNewDeck(ctx *gin.Context, req deck.CreateNewDeckRequestDTO) (dao.Deck, error) {
	deckID := s.resource.Helper.CreateNewUUID()
	cards := s.resource.DefaultCards

	if len(req.Cards) > 0 {
		filterCards := strings.Split(req.Cards[0], ",")
		cards.FilterCards(filterCards)
	}

	newDeck := dao.Deck{
		DeckID:    deckID,
		Remaining: uint(len(cards.Cards)),
		IsUsed:    false,
		Cards:     cards.ToCardDecks(deckID),
	}

	if req.IsShuffled {
		newDeck.Shuffle(cards)
	}

	if err := s.repository.StoreDeck(ctx, newDeck); err != nil {
		return dao.Deck{}, err
	}

	return newDeck, nil
}

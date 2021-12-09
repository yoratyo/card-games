package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/model/dao"
	"github.com/yoratyo/card-games/model/dto/deck"
	"gorm.io/gorm"
)

func (s *ServiceSuite) TestService_DrawCards() {
	s.Run("when error PatchDrawCardDeck", func() {
		deckID := uuid.New()
		req := deck.DrawCardsFromDeckRequestDTO{}

		expectedErr := errors.New("not found")

		s.repository.EXPECT().PatchDrawCardDeck(
			s.ctx,
			deckID,
			1,
			map[string]interface{}{"is_drawn": true},
		).Return([]dao.CardDeck{}, expectedErr)

		result, err := s.service.DrawCards(s.ctx, deckID, req)

		s.Equal(result, deck.DrawCardsFromDeckResponseDTO{})
		s.NotNil(err)
		s.Equal(err, expectedErr)
	})

	s.Run("when error count param is more than remaining deck", func() {
		deckID := uuid.New()
		req := deck.DrawCardsFromDeckRequestDTO{}

		expectedErr := errors.New("count param is more than remaining deck")

		s.repository.EXPECT().PatchDrawCardDeck(
			s.ctx,
			deckID,
			1,
			map[string]interface{}{"is_drawn": true},
		).Return([]dao.CardDeck{}, nil)

		result, err := s.service.DrawCards(s.ctx, deckID, req)

		s.Equal(result, deck.DrawCardsFromDeckResponseDTO{})
		s.NotNil(err)
		s.Equal(err, expectedErr)
	})

	s.Run("when error PatchDeckByID", func() {
		deckID := uuid.New()
		req := deck.DrawCardsFromDeckRequestDTO{
			Count: 1,
		}

		cardDecks := []dao.CardDeck{
			{
				DeckID:   deckID,
				Code:     "AS",
				Sequence: 5,
				Value:    "ACE",
				Suit:     "SPADES",
				IsDrawn:  false,
			},
		}

		expectedErr := errors.New("not found")

		s.repository.EXPECT().PatchDrawCardDeck(
			s.ctx,
			deckID,
			1,
			map[string]interface{}{"is_drawn": true},
		).Return(cardDecks, nil)
		s.repository.EXPECT().PatchDeckByID(
			s.ctx,
			deckID,
			map[string]interface{}{
				"remaining": gorm.Expr("remaining - ?", 1),
				"is_used":   true,
			},
		).Return(int64(0), expectedErr)

		result, err := s.service.DrawCards(s.ctx, deckID, req)

		s.Equal(result, deck.DrawCardsFromDeckResponseDTO{})
		s.NotNil(err)
		s.Equal(err, expectedErr)
	})

	s.Run("when error failed to update deck", func() {
		deckID := uuid.New()
		req := deck.DrawCardsFromDeckRequestDTO{
			Count: 1,
		}

		cardDecks := []dao.CardDeck{
			{
				DeckID:   deckID,
				Code:     "AS",
				Sequence: 5,
				Value:    "ACE",
				Suit:     "SPADES",
				IsDrawn:  false,
			},
		}

		expectedErr := errors.New("failed to update remaining deck")

		s.repository.EXPECT().PatchDrawCardDeck(
			s.ctx,
			deckID,
			1,
			map[string]interface{}{"is_drawn": true},
		).Return(cardDecks, nil)
		s.repository.EXPECT().PatchDeckByID(
			s.ctx,
			deckID,
			map[string]interface{}{
				"remaining": gorm.Expr("remaining - ?", 1),
				"is_used":   true,
			},
		).Return(int64(0), nil)

		result, err := s.service.DrawCards(s.ctx, deckID, req)

		s.Equal(result, deck.DrawCardsFromDeckResponseDTO{})
		s.NotNil(err)
		s.Equal(err, expectedErr)
	})

	s.Run("when success", func() {
		deckID := uuid.New()
		req := deck.DrawCardsFromDeckRequestDTO{
			Count: 1,
		}

		cardDecks := []dao.CardDeck{
			{
				DeckID:   deckID,
				Code:     "AS",
				Sequence: 5,
				Value:    "ACE",
				Suit:     "SPADES",
				IsDrawn:  false,
			},
		}

		expectedResult := deck.DrawCardsFromDeckResponseDTO{
			ListCard: dao.ListCard{Cards: []dao.Card{
				{
					Value: cardDecks[0].Value,
					Suit:  cardDecks[0].Suit,
					Code:  cardDecks[0].Code,
				},
			}},
		}

		s.repository.EXPECT().PatchDrawCardDeck(
			s.ctx,
			deckID,
			1,
			map[string]interface{}{"is_drawn": true},
		).Return(cardDecks, nil)
		s.repository.EXPECT().PatchDeckByID(
			s.ctx,
			deckID,
			map[string]interface{}{
				"remaining": gorm.Expr("remaining - ?", 1),
				"is_used":   true,
			},
		).Return(int64(1), nil)

		result, err := s.service.DrawCards(s.ctx, deckID, req)

		s.Equal(result, expectedResult)
		s.Nil(err)
	})
}

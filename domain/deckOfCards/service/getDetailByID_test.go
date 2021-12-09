package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/model/dao"
	"github.com/yoratyo/card-games/model/dto/deck"
)

func (s *ServiceSuite) TestService_GetDetailByID() {
	s.Run("when error GetDeckDetailByID", func() {
		deckID := uuid.New()

		expectedErr := errors.New("not found")

		s.repository.EXPECT().GetDeckDetailByID(s.ctx, deckID).Return(dao.Deck{}, expectedErr)

		result, err := s.service.GetDetailByID(s.ctx, deckID)

		s.Equal(result, deck.OpenDeckResponseDTO{})
		s.NotNil(err)
		s.Equal(err, expectedErr)
	})

	s.Run("when error GetCardDeckByDeckID", func() {
		deckID := uuid.New()

		expectedErr := errors.New("not found")

		deckDetail := dao.Deck{
			DeckID: deckID,
		}

		s.repository.EXPECT().GetDeckDetailByID(s.ctx, deckID).Return(deckDetail, nil)
		s.repository.EXPECT().GetCardDeckByDeckID(s.ctx, deckID).Return([]dao.CardDeck{}, expectedErr)

		result, err := s.service.GetDetailByID(s.ctx, deckID)

		s.Equal(result, deck.OpenDeckResponseDTO{})
		s.NotNil(err)
		s.Equal(err, expectedErr)
	})

	s.Run("when success", func() {
		deckID := uuid.New()

		deckDetail := dao.Deck{
			DeckID:     deckID,
			IsShuffled: true,
			Remaining:  1,
			IsUsed:     true,
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

		expectedResult := deck.OpenDeckResponseDTO{
			DeckID:    deckDetail.DeckID,
			Shuffled:  deckDetail.IsShuffled,
			Remaining: deckDetail.Remaining,
			ListCard: dao.ListCard{Cards: []dao.Card{
				{
					Value: cardDecks[0].Value,
					Suit:  cardDecks[0].Suit,
					Code:  cardDecks[0].Code,
				},
			}},
		}

		s.repository.EXPECT().GetDeckDetailByID(s.ctx, deckID).Return(deckDetail, nil)
		s.repository.EXPECT().GetCardDeckByDeckID(s.ctx, deckID).Return(cardDecks, nil)

		result, err := s.service.GetDetailByID(s.ctx, deckID)

		s.Equal(result, expectedResult)
		s.Nil(err)
	})
}

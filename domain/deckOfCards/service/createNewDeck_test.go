package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/model/dao"
	"github.com/yoratyo/card-games/model/dto/deck"
)

func (s *ServiceSuite) TestService_CreateNewDeck() {
	s.Run("when error StoreDeck", func() {
		uuidStr := "86c3cb8b-f100-44f0-b274-bdfcff71a561"
		uuid, _ := uuid.Parse(uuidStr)

		s.helper.EXPECT().CreateNewUUID().Return(uuid)

		req := deck.CreateNewDeckRequestDTO{
			IsShuffled: true,
			Cards:      []string{"AS"},
		}

		newDeck := dao.Deck{
			DeckID:     uuid,
			IsShuffled: true,
			Remaining:  1,
			IsUsed:     false,
			Cards: []dao.CardDeck{
				{
					DeckID:   uuid,
					Code:     "AS",
					Sequence: 1,
					Value:    "ACE",
					Suit:     "SPADES",
					IsDrawn:  false,
				},
			},
		}

		expectedErr := errors.New("duplicate data")

		s.repository.EXPECT().StoreDeck(s.ctx, newDeck).Return(expectedErr)

		result, err := s.service.CreateNewDeck(s.ctx, req)

		s.Equal(result, dao.Deck{})
		s.NotNil(err)
		s.Equal(err, expectedErr)
	})

	s.Run("when success", func() {
		uuidStr := "86c3cb8b-f100-44f0-b274-bdfcff71a561"
		uuid, _ := uuid.Parse(uuidStr)

		s.helper.EXPECT().CreateNewUUID().Return(uuid)

		req := deck.CreateNewDeckRequestDTO{
			IsShuffled: true,
			Cards:      []string{"AS"},
		}

		newDeck := dao.Deck{
			DeckID:     uuid,
			IsShuffled: true,
			Remaining:  1,
			IsUsed:     false,
			Cards: []dao.CardDeck{
				{
					DeckID:   uuid,
					Code:     "AS",
					Sequence: 1,
					Value:    "ACE",
					Suit:     "SPADES",
					IsDrawn:  false,
				},
			},
		}

		s.repository.EXPECT().StoreDeck(s.ctx, newDeck).Return(nil)

		result, err := s.service.CreateNewDeck(s.ctx, req)

		s.Equal(result, newDeck)
		s.Nil(err)
	})
}

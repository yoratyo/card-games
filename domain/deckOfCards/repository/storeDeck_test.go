package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/config/database"
	"github.com/yoratyo/card-games/model/dao"
)

func (s *RepositorySuite) TestRepository_StoreDeck() {
	s.Run("when error get transaction", func() {
		deckID := uuid.New()
		request := dao.Deck{
			DeckID:    deckID,
			Remaining: 0,
			IsUsed:    true,
		}
		expectedError := errors.New("invalid transaction")
		s.ctx.Set(database.DB_TRANSACTION, "wrong trx")

		err := s.repository.StoreDeck(s.ctx, request)
		s.NotNil(err)
		s.Equal(err, expectedError)
	})

	s.Run("when error create", func() {
		deckID := uuid.New()
		request := dao.Deck{
			DeckID:    deckID,
			Remaining: 0,
			IsUsed:    true,
		}
		expectedError := errors.New("duplicate entry")
		s.ctx.Set(database.DB_TRANSACTION, s.database)

		s.database.EXPECT().WithContext(s.ctx).Return(s.database)
		s.database.EXPECT().Model(&dao.Deck{}).Return(s.database)
		s.database.EXPECT().Create(&request).Return(s.database)
		s.database.EXPECT().Error().Return(expectedError)

		err := s.repository.StoreDeck(s.ctx, request)
		s.NotNil(err)
		s.Equal(err, expectedError)
	})

	s.Run("when success", func() {
		deckID := uuid.New()
		request := dao.Deck{
			DeckID:    deckID,
			Remaining: 0,
			IsUsed:    true,
		}

		s.database.EXPECT().WithContext(s.ctx).Return(s.database)
		s.database.EXPECT().Model(&dao.Deck{}).Return(s.database)
		s.database.EXPECT().Create(&request).Return(s.database)
		s.database.EXPECT().Error().Return(nil)

		err := s.repository.StoreDeck(s.ctx, request)
		s.Nil(err)
	})
}

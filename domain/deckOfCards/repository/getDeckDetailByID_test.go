package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/config/database"
	"github.com/yoratyo/card-games/model/dao"
)

func (s *RepositorySuite) TestRepository_GetDeckDetailByID() {
	s.Run("when error get transaction", func() {
		s.ctx.Set(database.DB_TRANSACTION, "wrong trx")
		deckID := uuid.New()
		expectedError := errors.New("invalid transaction")

		result, err := s.repository.GetDeckDetailByID(s.ctx, deckID)
		s.Equal(result, dao.Deck{})
		s.NotNil(err)
		s.Equal(err, expectedError)
	})

	s.Run("when error Take", func() {
		s.ctx.Set(database.DB_TRANSACTION, s.database)
		deckID := uuid.New()
		expectedError := errors.New("not found")

		var item dao.Deck
		s.database.EXPECT().WithContext(s.ctx).Return(s.database)
		s.database.EXPECT().Model(&dao.Deck{}).Return(s.database)
		s.database.EXPECT().Where("deck_id = ?", deckID).Return(s.database)
		s.database.EXPECT().Take(&item).Return(s.database)
		s.database.EXPECT().Error().Return(expectedError)

		result, err := s.repository.GetDeckDetailByID(s.ctx, deckID)
		s.Equal(result, dao.Deck{})
		s.NotNil(err)
		s.Equal(err, expectedError)
	})

	s.Run("when success", func() {
		deckID := uuid.New()
		expectedResult := dao.Deck{
			DeckID:    deckID,
			Remaining: 1,
			IsUsed:    true,
		}

		var item dao.Deck
		s.database.EXPECT().WithContext(s.ctx).Return(s.database)
		s.database.EXPECT().Model(&dao.Deck{}).Return(s.database)
		s.database.EXPECT().Where("deck_id = ?", deckID).Return(s.database)
		s.database.EXPECT().Take(&item).SetArg(0, expectedResult).Return(s.database)
		s.database.EXPECT().Error().Return(nil)

		result, err := s.repository.GetDeckDetailByID(s.ctx, deckID)
		s.Equal(result, expectedResult)
		s.Nil(err)
	})
}

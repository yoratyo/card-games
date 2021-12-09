package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/config/database"
	"github.com/yoratyo/card-games/model/dao"
)

func (s *RepositorySuite) TestRepository_PatchDeckByID() {
	s.Run("when error get transaction", func() {
		s.ctx.Set(database.DB_TRANSACTION, "wrong trx")
		deckID := uuid.New()
		data := map[string]interface{}{
			"is_used": true,
		}
		expectedError := errors.New("invalid transaction")

		result, err := s.repository.PatchDeckByID(s.ctx, deckID, data)
		s.Equal(result, int64(0))
		s.NotNil(err)
		s.Equal(err, expectedError)
	})

	s.Run("when error update", func() {
		s.ctx.Set(database.DB_TRANSACTION, s.database)
		deckID := uuid.New()
		data := map[string]interface{}{
			"is_used": true,
		}
		expectedError := errors.New("bad connection")

		s.database.EXPECT().WithContext(s.ctx).Return(s.database)
		s.database.EXPECT().Model(&dao.Deck{}).Return(s.database)
		s.database.EXPECT().Where("deck_id = ?", deckID).Return(s.database)
		s.database.EXPECT().Updates(data).Return(s.database)
		s.database.EXPECT().Error().Return(expectedError)

		result, err := s.repository.PatchDeckByID(s.ctx, deckID, data)
		s.Equal(result, int64(0))
		s.NotNil(err)
		s.Equal(err, expectedError)
	})

	s.Run("when success", func() {
		deckID := uuid.New()
		data := map[string]interface{}{
			"is_used": true,
		}

		s.database.EXPECT().WithContext(s.ctx).Return(s.database)
		s.database.EXPECT().Model(&dao.Deck{}).Return(s.database)
		s.database.EXPECT().Where("deck_id = ?", deckID).Return(s.database)
		s.database.EXPECT().Updates(data).Return(s.database)
		s.database.EXPECT().Error().Return(nil)
		s.database.EXPECT().RowsAffected().Return(int64(1))

		result, err := s.repository.PatchDeckByID(s.ctx, deckID, data)
		s.Equal(result, int64(1))
		s.Nil(err)
	})
}

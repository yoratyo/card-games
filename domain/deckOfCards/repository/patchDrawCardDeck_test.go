package repository

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/config/database"
	"github.com/yoratyo/card-games/model/dao"
)

func (s *RepositorySuite) TestRepository_PatchDrawCardDeck() {
	s.Run("when error get transaction", func() {
		s.ctx.Set(database.DB_TRANSACTION, "wrong trx")
		deckID := uuid.New()
		data := map[string]interface{}{
			"is_drawn": true,
		}
		expectedError := errors.New("invalid transaction")

		result, err := s.repository.PatchDrawCardDeck(s.ctx, deckID, 1, data)
		s.Equal(result, []dao.CardDeck{})
		s.NotNil(err)
		s.Equal(err, expectedError)
	})

	s.Run("when error find", func() {
		s.ctx.Set(database.DB_TRANSACTION, s.database)
		deckID := uuid.New()
		data := map[string]interface{}{
			"is_drawn": true,
		}
		count := 1
		expectedError := errors.New("bad connection")

		var items []dao.CardDeck

		s.database.EXPECT().WithContext(s.ctx).Return(s.database)
		s.database.EXPECT().Model(&dao.CardDeck{}).Return(s.database)
		s.database.EXPECT().Where("deck_id = ?", deckID).Return(s.database)
		s.database.EXPECT().Where("is_drawn = ?", false).Return(s.database)
		s.database.EXPECT().Order("sequence").Return(s.database)
		s.database.EXPECT().Limit(count).Return(s.database)
		s.database.EXPECT().Find(&items).Return(s.database)
		s.database.EXPECT().Error().Return(expectedError)

		result, err := s.repository.PatchDrawCardDeck(s.ctx, deckID, count, data)
		s.Equal(result, []dao.CardDeck{})
		s.NotNil(err)
		s.Equal(err, expectedError)
	})

	s.Run("when error not found record", func() {
		deckID := uuid.New()
		data := map[string]interface{}{
			"is_drawn": true,
		}
		count := 1
		expectedError := errors.New("deck is empty or invalid")

		var items []dao.CardDeck

		s.database.EXPECT().WithContext(s.ctx).Return(s.database)
		s.database.EXPECT().Model(&dao.CardDeck{}).Return(s.database)
		s.database.EXPECT().Where("deck_id = ?", deckID).Return(s.database)
		s.database.EXPECT().Where("is_drawn = ?", false).Return(s.database)
		s.database.EXPECT().Order("sequence").Return(s.database)
		s.database.EXPECT().Limit(count).Return(s.database)
		s.database.EXPECT().Find(&items).Return(s.database)
		s.database.EXPECT().Error().Return(nil)

		result, err := s.repository.PatchDrawCardDeck(s.ctx, deckID, count, data)
		s.Equal(result, []dao.CardDeck{})
		s.NotNil(err)
		s.Equal(err, expectedError)
	})

	s.Run("when error update", func() {
		deckID := uuid.New()
		data := map[string]interface{}{
			"is_drawn": true,
		}
		count := 1
		expectedError := errors.New("not found")
		expectedResult := []dao.CardDeck{
			{
				DeckID:   deckID,
				Code:     "AS",
				Sequence: 5,
				Value:    "ACE",
				Suit:     "SPADES",
				IsDrawn:  false,
			},
		}

		var items []dao.CardDeck

		s.database.EXPECT().WithContext(s.ctx).Return(s.database)
		s.database.EXPECT().Model(&dao.CardDeck{}).Return(s.database)
		s.database.EXPECT().Where("deck_id = ?", deckID).Return(s.database)
		s.database.EXPECT().Where("is_drawn = ?", false).Return(s.database)
		s.database.EXPECT().Order("sequence").Return(s.database)
		s.database.EXPECT().Limit(count).Return(s.database)
		s.database.EXPECT().Find(&items).SetArg(0, expectedResult).Return(s.database)

		gomock.InOrder(
			s.database.EXPECT().Error().Return(nil),
			s.database.EXPECT().Updates(data).Return(s.database),
			s.database.EXPECT().Error().Return(expectedError),
		)

		result, err := s.repository.PatchDrawCardDeck(s.ctx, deckID, count, data)
		s.Equal(result, []dao.CardDeck{})
		s.NotNil(err)
		s.Equal(err, expectedError)
	})

	s.Run("when success", func() {
		deckID := uuid.New()
		data := map[string]interface{}{
			"is_drawn": true,
		}
		count := 1
		expectedResult := []dao.CardDeck{
			{
				DeckID:   deckID,
				Code:     "AS",
				Sequence: 5,
				Value:    "ACE",
				Suit:     "SPADES",
				IsDrawn:  false,
			},
		}

		var items []dao.CardDeck

		s.database.EXPECT().WithContext(s.ctx).Return(s.database)
		s.database.EXPECT().Model(&dao.CardDeck{}).Return(s.database)
		s.database.EXPECT().Where("deck_id = ?", deckID).Return(s.database)
		s.database.EXPECT().Where("is_drawn = ?", false).Return(s.database)
		s.database.EXPECT().Order("sequence").Return(s.database)
		s.database.EXPECT().Limit(count).Return(s.database)
		s.database.EXPECT().Find(&items).SetArg(0, expectedResult).Return(s.database)
		s.database.EXPECT().Error().Return(nil).AnyTimes()
		s.database.EXPECT().Updates(data).Return(s.database)

		result, err := s.repository.PatchDrawCardDeck(s.ctx, deckID, count, data)
		s.Equal(result, expectedResult)
		s.Nil(err)
	})
}

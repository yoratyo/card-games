package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/config/database"
	"github.com/yoratyo/card-games/model/dao"
)

func (s *RepositorySuite) TestRepository_GetCardDeckByDeckID() {
	s.Run("when error get transaction", func() {
		s.ctx.Set(database.DB_TRANSACTION, "wrong trx")
		deckID := uuid.New()
		expectedError := errors.New("invalid transaction")

		result, err := s.repository.GetCardDeckByDeckID(s.ctx, deckID)
		s.Equal(result, []dao.CardDeck{})
		s.NotNil(err)
		s.Equal(err, expectedError)
	})

	s.Run("when error Take", func() {
		s.ctx.Set(database.DB_TRANSACTION, s.database)
		deckID := uuid.New()
		expectedError := errors.New("not found")

		var items []dao.CardDeck
		s.database.EXPECT().WithContext(s.ctx).Return(s.database)
		s.database.EXPECT().Model(&dao.CardDeck{}).Return(s.database)
		s.database.EXPECT().Where("deck_id = ?", deckID).Return(s.database)
		s.database.EXPECT().Where("is_drawn = ?", false).Return(s.database)
		s.database.EXPECT().Order("sequence").Return(s.database)
		s.database.EXPECT().Find(&items).Return(s.database)
		s.database.EXPECT().Error().Return(expectedError)

		result, err := s.repository.GetCardDeckByDeckID(s.ctx, deckID)
		s.Equal(result, []dao.CardDeck{})
		s.NotNil(err)
		s.Equal(err, expectedError)
	})

	s.Run("when success", func() {
		deckID := uuid.New()
		expectedResult := []dao.CardDeck{
			{
				DeckID:   deckID,
				Code:     "AS",
				Sequence: 1,
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
		s.database.EXPECT().Find(&items).SetArg(0, expectedResult).Return(s.database)
		s.database.EXPECT().Error().Return(nil)

		result, err := s.repository.GetCardDeckByDeckID(s.ctx, deckID)
		s.Equal(result, expectedResult)
		s.Nil(err)
	})
}

package deck

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/model/dao"
	"github.com/yoratyo/card-games/model/dto/deck"
	"net/http"
	"net/http/httptest"
	"net/url"
)

func (s *HandlerSuite) TestHandler_DrawCardsFromDeck() {
	s.Run("when error parse uuid", func() {
		s.recorder = httptest.NewRecorder()
		c, _ := gin.CreateTestContext(s.recorder)
		s.ctx = c

		s.ctx.Params = []gin.Param{{Key: "deck_id", Value: "v"}}
		expectedError := errors.New("invalid UUID length: 1")

		s.handler.DrawCardsFromDeck(s.ctx)

		var response map[string]string
		err := json.Unmarshal([]byte(s.recorder.Body.String()), &response)
		res, exists := response["message"]

		s.Nil(err)
		s.True(exists)
		s.Equal(expectedError.Error(), res)
		s.Equal(s.recorder.Code, http.StatusBadRequest)
	})

	s.Run("when error binding param", func() {
		s.recorder = httptest.NewRecorder()
		c, _ := gin.CreateTestContext(s.recorder)
		s.ctx = c
		s.ctx.Request = &http.Request{
			Header: make(http.Header),
			Form: url.Values{
				"count": []string{"x"},
			},
		}
		s.ctx.Params = []gin.Param{
			{Key: "deck_id", Value: "cd0693dc-4d3c-428a-b1aa-d1e4408d1c47"},
		}

		expectedError := errors.New("strconv.ParseUint: parsing \"x\": invalid syntax")

		s.handler.DrawCardsFromDeck(s.ctx)

		var response map[string]string
		err := json.Unmarshal([]byte(s.recorder.Body.String()), &response)
		res, exists := response["message"]

		s.Nil(err)
		s.True(exists)
		s.Equal(expectedError.Error(), res)
		s.Equal(s.recorder.Code, http.StatusBadRequest)
	})

	s.Run("when error call domain", func() {
		s.recorder = httptest.NewRecorder()
		c, _ := gin.CreateTestContext(s.recorder)
		s.ctx = c
		s.ctx.Request = &http.Request{
			Header: make(http.Header),
			Form: url.Values{
				"count": []string{"1"},
			},
		}
		ID := "cd0693dc-4d3c-428a-b1aa-d1e4408d1c47"
		s.ctx.Params = []gin.Param{
			{Key: "deck_id", Value: ID},
		}

		s.ctx.Params = []gin.Param{
			{Key: "deck_id", Value: ID},
		}
		deckID, _ := uuid.Parse(ID)

		params := deck.DrawCardsFromDeckRequestDTO{Count: 1}

		expectedError := errors.New("deck is empty")

		s.database.EXPECT().Begin().Return(s.database)
		s.service.EXPECT().DrawCards(s.ctx, deckID, params).Return(deck.DrawCardsFromDeckResponseDTO{}, expectedError)
		s.database.EXPECT().Rollback().Return(s.database)

		s.handler.DrawCardsFromDeck(s.ctx)

		var response map[string]string
		err := json.Unmarshal([]byte(s.recorder.Body.String()), &response)
		res, exists := response["message"]

		s.Nil(err)
		s.True(exists)
		s.Equal(expectedError.Error(), res)
		s.Equal(s.recorder.Code, http.StatusBadRequest)
	})

	s.Run("when error commit transaction", func() {
		s.recorder = httptest.NewRecorder()
		c, _ := gin.CreateTestContext(s.recorder)
		s.ctx = c
		s.ctx.Request = &http.Request{
			Header: make(http.Header),
			Form: url.Values{
				"count": []string{"1"},
			},
		}
		ID := "cd0693dc-4d3c-428a-b1aa-d1e4408d1c47"
		s.ctx.Params = []gin.Param{
			{Key: "deck_id", Value: ID},
		}

		s.ctx.Params = []gin.Param{
			{Key: "deck_id", Value: ID},
		}
		deckID, _ := uuid.Parse(ID)

		params := deck.DrawCardsFromDeckRequestDTO{Count: 1}

		expectedError := errors.New("failed to commit")

		s.database.EXPECT().Begin().Return(s.database)
		s.service.EXPECT().DrawCards(s.ctx, deckID, params).Return(deck.DrawCardsFromDeckResponseDTO{}, nil)
		s.database.EXPECT().Rollback().Return(s.database)
		s.database.EXPECT().Commit().Return(s.database)
		s.database.EXPECT().Error().Return(expectedError)

		s.handler.DrawCardsFromDeck(s.ctx)

		var response map[string]string
		err := json.Unmarshal([]byte(s.recorder.Body.String()), &response)
		res, exists := response["message"]

		s.Nil(err)
		s.True(exists)
		s.Equal(expectedError.Error(), res)
		s.Equal(s.recorder.Code, http.StatusBadRequest)
	})

	s.Run("when success", func() {
		s.recorder = httptest.NewRecorder()
		c, _ := gin.CreateTestContext(s.recorder)
		s.ctx = c
		s.ctx.Request = &http.Request{
			Header: make(http.Header),
			Form: url.Values{
				"count": []string{"1"},
			},
		}
		ID := "cd0693dc-4d3c-428a-b1aa-d1e4408d1c47"
		s.ctx.Params = []gin.Param{
			{Key: "deck_id", Value: ID},
		}

		s.ctx.Params = []gin.Param{
			{Key: "deck_id", Value: ID},
		}
		deckID, _ := uuid.Parse(ID)

		params := deck.DrawCardsFromDeckRequestDTO{Count: 1}

		expectedResult := deck.DrawCardsFromDeckResponseDTO{
			ListCard: dao.ListCard{Cards: []dao.Card{
				{
					Value: "4",
					Suit:  "SPADES",
					Code:  "4S",
				},
			},
			},
		}

		s.database.EXPECT().Begin().Return(s.database)
		s.service.EXPECT().DrawCards(s.ctx, deckID, params).Return(expectedResult, nil)
		s.database.EXPECT().Commit().Return(s.database)
		s.database.EXPECT().Error().Return(nil)

		s.handler.DrawCardsFromDeck(s.ctx)

		var response deck.DrawCardsFromDeckResponseDTO
		json.Unmarshal([]byte(s.recorder.Body.String()), &response)

		s.Equal(expectedResult, response)
		s.Equal(s.recorder.Code, http.StatusOK)
	})
}

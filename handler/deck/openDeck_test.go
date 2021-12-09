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
)

func (s *HandlerSuite) TestHandler_OpenDeck() {
	s.Run("when error parse uuid", func() {
		s.recorder = httptest.NewRecorder()
		c, _ := gin.CreateTestContext(s.recorder)
		s.ctx = c

		s.ctx.Params = []gin.Param{{Key: "deck_id", Value: "v"}}
		expectedError := errors.New("invalid UUID length: 1")

		s.handler.OpenDeck(s.ctx)

		var response map[string]string
		err := json.Unmarshal([]byte(s.recorder.Body.String()), &response)
		res, exists := response["message"]

		s.Nil(err)
		s.True(exists)
		s.Equal(expectedError.Error(), res)
		s.Equal(s.recorder.Code, http.StatusBadRequest)
	})

	s.Run("when error GetDetailByID", func() {
		s.recorder = httptest.NewRecorder()
		c, _ := gin.CreateTestContext(s.recorder)
		s.ctx = c

		ID := "cd0693dc-4d3c-428a-b1aa-d1e4408d1c47"
		s.ctx.Params = []gin.Param{
			{Key: "deck_id", Value: ID},
		}
		deckID, _ := uuid.Parse(ID)
		expectedError := errors.New("not found")

		s.service.EXPECT().GetDetailByID(s.ctx, deckID).Return(deck.OpenDeckResponseDTO{}, expectedError)

		s.handler.OpenDeck(s.ctx)

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

		ID := "cd0693dc-4d3c-428a-b1aa-d1e4408d1c47"
		s.ctx.Params = []gin.Param{
			{Key: "deck_id", Value: ID},
		}
		deckID, _ := uuid.Parse(ID)
		expectedResult := deck.OpenDeckResponseDTO{
			DeckID:    deckID,
			Shuffled:  true,
			Remaining: 1,
			ListCard: dao.ListCard{Cards: []dao.Card{
				{
					Value: "4",
					Suit:  "SPADES",
					Code:  "4S",
				},
			}},
		}

		s.service.EXPECT().GetDetailByID(s.ctx, deckID).Return(expectedResult, nil)

		s.handler.OpenDeck(s.ctx)

		var response deck.OpenDeckResponseDTO
		json.Unmarshal([]byte(s.recorder.Body.String()), &response)

		s.Equal(expectedResult, response)
		s.Equal(s.recorder.Code, http.StatusOK)
	})
}

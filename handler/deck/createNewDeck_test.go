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

func (s *HandlerSuite) TestHandler_CreateNewDeck() {
	s.Run("when error binding param", func() {
		s.recorder = httptest.NewRecorder()
		c, _ := gin.CreateTestContext(s.recorder)
		s.ctx = c
		s.ctx.Request = &http.Request{
			Header: make(http.Header),
			Form: url.Values{
				"shuffle": []string{"x"},
			},
		}

		expectedError := errors.New("strconv.ParseBool: parsing \"x\": invalid syntax")

		s.handler.CreateNewDeck(s.ctx)

		var response map[string]string
		err := json.Unmarshal([]byte(s.recorder.Body.String()), &response)
		res, exists := response["message"]

		s.Nil(err)
		s.True(exists)
		s.Equal(expectedError.Error(), res)
		s.Equal(s.recorder.Code, http.StatusBadRequest)
	})

	s.Run("when error call service domain", func() {
		s.recorder = httptest.NewRecorder()
		c, _ := gin.CreateTestContext(s.recorder)
		s.ctx = c
		s.ctx.Request = &http.Request{
			Header: make(http.Header),
			Form: url.Values{
				"shuffle": []string{"true"},
			},
		}

		params := deck.CreateNewDeckRequestDTO{
			IsShuffled: true,
		}

		expectedError := errors.New("duplicate entry")

		s.database.EXPECT().Begin().Return(s.database)
		s.service.EXPECT().CreateNewDeck(s.ctx, params).Return(dao.Deck{}, expectedError)
		s.database.EXPECT().Rollback().Return(s.database)

		s.handler.CreateNewDeck(s.ctx)

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
				"shuffle": []string{"true"},
			},
		}

		params := deck.CreateNewDeckRequestDTO{
			IsShuffled: true,
		}

		expectedError := errors.New("failed to commit")

		s.database.EXPECT().Begin().Return(s.database)
		s.service.EXPECT().CreateNewDeck(s.ctx, params).Return(dao.Deck{}, nil)
		s.database.EXPECT().Commit().Return(s.database)
		s.database.EXPECT().Error().Return(expectedError)
		s.database.EXPECT().Rollback().Return(s.database)

		s.handler.CreateNewDeck(s.ctx)

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
				"shuffle": []string{"true"},
			},
		}

		params := deck.CreateNewDeckRequestDTO{
			IsShuffled: true,
		}

		deckID := uuid.New()
		newDeck := dao.Deck{
			DeckID:     deckID,
			IsShuffled: true,
			Remaining:  1,
		}

		s.database.EXPECT().Begin().Return(s.database)
		s.service.EXPECT().CreateNewDeck(s.ctx, params).Return(newDeck, nil)
		s.database.EXPECT().Commit().Return(s.database)
		s.database.EXPECT().Error().Return(nil)

		s.handler.CreateNewDeck(s.ctx)

		expectedResult := deck.CreateNewDeckResponseDTO{
			DeckID:    newDeck.DeckID,
			Shuffled:  newDeck.IsShuffled,
			Remaining: newDeck.Remaining,
		}

		var response deck.CreateNewDeckResponseDTO
		json.Unmarshal([]byte(s.recorder.Body.String()), &response)

		s.Equal(expectedResult, response)
		s.Equal(s.recorder.Code, http.StatusOK)
	})
}

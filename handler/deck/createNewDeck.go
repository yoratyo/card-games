package deck

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/card-games/config/database"
	"github.com/yoratyo/card-games/model/dto/deck"
	"net/http"
)

func (h handler) CreateNewDeck(c *gin.Context) {
	// get request param
	var params deck.CreateNewDeckRequestDTO
	err := c.Bind(&params)
	if err != nil {
		addError(c, err)
		return
	}

	// start transaction
	tx := h.resource.DB.Begin()
	c.Set(database.DB_TRANSACTION, tx)

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			addError(c, err)
			return
		}
	}()

	// call service domain
	newDeck, err := h.domain.DeckOfCards.CreateNewDeck(c, params)
	if err != nil {
		tx.Rollback()
		addError(c, err)

		return
	}

	// commit transaction
	if err := tx.Commit().Error(); err != nil {
		tx.Rollback()
		addError(c, err)

		return
	}

	c.JSON(http.StatusOK, deck.CreateNewDeckResponseDTO{
		DeckID:    newDeck.DeckID,
		Shuffled:  newDeck.IsShuffled,
		Remaining: newDeck.Remaining,
	})
}

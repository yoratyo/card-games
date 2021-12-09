package deck

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/config/database"
	"github.com/yoratyo/card-games/model/dto/deck"
	"net/http"
)

func (h handler) DrawCardsFromDeck(c *gin.Context) {
	var err error

	deckID, err := uuid.Parse(c.Params.ByName("deck_id"))
	if err != nil {
		addError(c, err)

		return
	}

	// get request param
	var params deck.DrawCardsFromDeckRequestDTO
	err = c.Bind(&params)
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
	drawnCards, err := h.domain.DeckOfCards.DrawCards(c, deckID, params)
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

	c.JSON(http.StatusOK, drawnCards)
}

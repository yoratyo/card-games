package deck

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (h handler) OpenDeck(c *gin.Context) {
	deckID, err := uuid.Parse(c.Params.ByName("deck_id"))
	if err != nil {
		addError(c, err)

		return
	}

	newDeck, err := h.domain.DeckOfCards.GetDetailByID(c, deckID)
	if err != nil {
		addError(c, err)

		return
	}

	c.JSON(http.StatusOK, newDeck)
}

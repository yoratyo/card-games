package app

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(a *App) *gin.Engine {
	router := a.Engine

	// API v1
	v1 := router.Group("/v1")

	// Deck
	v1.POST("/deck", a.Handler.Deck.CreateNewDeck)
	v1.GET("/deck/:deck_id", a.Handler.Deck.OpenDeck)
	v1.PATCH("/deck/:deck_id/draw", a.Handler.Deck.DrawCardsFromDeck)

	return router
}

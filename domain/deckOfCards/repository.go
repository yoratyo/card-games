package deckOfCards

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/model/dao"
)

type Repository interface {
	StoreDeck(ctx *gin.Context, req dao.Deck) error
	GetDeckDetailByID(ctx *gin.Context, ID uuid.UUID) (dao.Deck, error)
	GetCardDeckByDeckID(ctx *gin.Context, deckID uuid.UUID) ([]dao.CardDeck, error)
	PatchDeckByID(ctx *gin.Context, ID uuid.UUID, data map[string]interface{}) (int64, error)
	PatchDrawCardDeck(ctx *gin.Context, deckID uuid.UUID, count int, data map[string]interface{}) ([]dao.CardDeck, error)
}

package deckOfCards

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/model/dao"
	"github.com/yoratyo/card-games/model/dto/deck"
)

type Service interface {
	CreateNewDeck(ctx *gin.Context, req deck.CreateNewDeckRequestDTO) (dao.Deck, error)
	GetDetailByID(ctx *gin.Context, ID uuid.UUID) (deck.OpenDeckResponseDTO, error)
	DrawCards(ctx *gin.Context, ID uuid.UUID, req deck.DrawCardsFromDeckRequestDTO) (deck.DrawCardsFromDeckResponseDTO, error)
}

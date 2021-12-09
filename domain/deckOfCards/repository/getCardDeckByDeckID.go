package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/config/database"
	"github.com/yoratyo/card-games/model/dao"
)

func (r *repository) GetCardDeckByDeckID(ctx *gin.Context, deckID uuid.UUID) ([]dao.CardDeck, error) {
	orm, err := database.GetORMTransaction(ctx, r.resource.DB)
	if err != nil {
		return []dao.CardDeck{}, err
	}

	var items []dao.CardDeck

	query := orm.
		WithContext(ctx).
		Model(&dao.CardDeck{}).
		Where("deck_id = ?", deckID).
		Where("is_drawn = ?", false).
		Order("sequence")

	if err := query.Find(&items).Error(); err != nil {
		return []dao.CardDeck{}, err
	}

	return items, nil
}

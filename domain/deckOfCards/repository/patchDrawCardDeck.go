package repository

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/config/database"
	"github.com/yoratyo/card-games/model/dao"
)

func (r *repository) PatchDrawCardDeck(
	ctx *gin.Context,
	deckID uuid.UUID,
	count int,
	data map[string]interface{},
) ([]dao.CardDeck, error) {
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
		Order("sequence").
		Limit(count)

	if err := query.Find(&items).Error(); err != nil {
		return []dao.CardDeck{}, err
	}

	if len(items) == 0 {
		return []dao.CardDeck{}, errors.New("deck is empty or invalid")
	}

	result := query.Updates(data)
	err = result.Error()
	if err != nil {
		return []dao.CardDeck{}, err
	}

	return items, nil
}

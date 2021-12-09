package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/config/database"
	"github.com/yoratyo/card-games/model/dao"
)

func (r *repository) GetDeckDetailByID(ctx *gin.Context, ID uuid.UUID) (dao.Deck, error) {
	orm, err := database.GetORMTransaction(ctx, r.resource.DB)
	if err != nil {
		return dao.Deck{}, err
	}

	var item dao.Deck

	query := orm.
		WithContext(ctx).
		Model(&dao.Deck{}).
		Where("deck_id = ?", ID)

	if err := query.Take(&item).Error(); err != nil {
		return dao.Deck{}, err
	}

	return item, nil
}

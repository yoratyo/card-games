package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/card-games/config/database"
	"github.com/yoratyo/card-games/model/dao"
)

func (r *repository) StoreDeck(ctx *gin.Context, req dao.Deck) error {
	orm, err := database.GetORMTransaction(ctx, r.resource.DB)
	if err != nil {
		return err
	}

	if err := orm.
		WithContext(ctx).
		Model(&dao.Deck{}).
		Create(&req).
		Error(); err != nil {
		return err
	}

	return nil
}

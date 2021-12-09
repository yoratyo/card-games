package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yoratyo/card-games/config/database"
	"github.com/yoratyo/card-games/model/dao"
)

func (r *repository) PatchDeckByID(ctx *gin.Context, ID uuid.UUID, data map[string]interface{}) (int64, error) {
	orm, err := database.GetORMTransaction(ctx, r.resource.DB)
	if err != nil {
		return 0, err
	}

	result := orm.
		WithContext(ctx).
		Model(&dao.Deck{}).
		Where("deck_id = ?", ID).
		Updates(data)
	err = result.Error()
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

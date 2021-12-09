package database

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/card-games/model/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

const DB_TRANSACTION = "DB-TRANSACTION"

func NewDB() (Gormw, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	dialectMySQL := mysql.Open(dsn)

	db, err := Openw(dialectMySQL, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&dao.Deck{}, &dao.CardDeck{})

	return db, err
}

func GetORMTransaction(c *gin.Context, db Gormw) (Gormw, error) {
	var ok bool

	trxInt, exist := c.Get(DB_TRANSACTION)
	if exist {
		if db, ok = trxInt.(Gormw); !ok {
			return nil, errors.New("invalid transaction")
		}
	}

	return db, nil
}

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/card-games/handler"
	"os"
)

type App struct {
	Engine  *gin.Engine
	Handler handler.Handler
}

func (a *App) Start() error {
	engine := NewRouter(a)

	return engine.Run(":" + os.Getenv("APP_PORT"))
}

func NewApp(handler handler.Handler) *App {
	return &App{
		Engine:  gin.Default(),
		Handler: handler,
	}
}

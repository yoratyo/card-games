package app

import (
	"github.com/yoratyo/card-games/config"
	"github.com/yoratyo/card-games/config/database"
	"github.com/yoratyo/card-games/domain"
	deckOfCardsRepository "github.com/yoratyo/card-games/domain/deckOfCards/repository"
	deckOfCardsService "github.com/yoratyo/card-games/domain/deckOfCards/service"
	"github.com/yoratyo/card-games/handler"
	"github.com/yoratyo/card-games/handler/deck"
	"go.uber.org/dig"
)

func BuildRuntime() (*App, error) {
	c := dig.New()
	servicesConstructors := []interface{}{
		// config
		database.NewDB,
		config.NewResource,
		config.NewHelper,
		// Domain, Repository & Service
		deckOfCardsRepository.New,
		deckOfCardsService.New,
		domain.New,
		// Handler
		deck.NewHandler,
		handler.New,
		// App
		NewApp,
	}

	for _, service := range servicesConstructors {
		if err := c.Provide(service); err != nil {
			return nil, err
		}
	}

	var result *App
	err := c.Invoke(func(a *App) {
		result = a
	})
	return result, err
}

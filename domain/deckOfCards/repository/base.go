package repository

import (
	"github.com/yoratyo/card-games/config"
	"github.com/yoratyo/card-games/domain/deckOfCards"
)

type repository struct {
	resource config.Resource
}

func New(resource config.Resource) (deckOfCards.Repository, error) {
	return &repository{
		resource: resource,
	}, nil
}

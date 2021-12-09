package service

import (
	"github.com/yoratyo/card-games/config"
	"github.com/yoratyo/card-games/domain/deckOfCards"
)

type service struct {
	repository deckOfCards.Repository
	resource   config.Resource
}

func New(repository deckOfCards.Repository, resource config.Resource) (deckOfCards.Service, error) {
	return &service{
		repository: repository,
		resource:   resource,
	}, nil
}

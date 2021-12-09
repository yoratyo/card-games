package config

import (
	"github.com/google/uuid"
)

type Helper interface {
	CreateNewUUID() uuid.UUID
}

type helper struct{}

func NewHelper() Helper {
	return &helper{}
}

func (r *helper) CreateNewUUID() uuid.UUID {
	return uuid.New()
}

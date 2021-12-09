package config

import (
	"fmt"
	"github.com/yoratyo/card-games/config/database"
	"github.com/yoratyo/card-games/model/dao"
	"strconv"
)

type (
	Resource struct {
		DB           database.Gormw
		DefaultCards dao.ListCard
		Helper       Helper
	}
)

// NewResource .
func NewResource(database database.Gormw, helper Helper) Resource {
	return Resource{
		DB:           database,
		DefaultCards: NewDefaultCards(),
		Helper:       helper,
	}
}

func NewDefaultCards() dao.ListCard {
	var defaultCards []dao.Card

	for _, suit := range dao.Suits {
		for _, value := range dao.Values {
			v := value[0:1]
			if _, err := strconv.Atoi(value); err == nil {
				v = value
			}
			card := dao.Card{
				Code:  fmt.Sprintf("%s%s", v, suit[0:1]),
				Value: value,
				Suit:  suit,
			}
			defaultCards = append(defaultCards, card)
		}
	}

	return dao.ListCard{Cards: defaultCards}
}

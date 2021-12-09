package deck

type (
	CreateNewDeckRequestDTO struct {
		IsShuffled bool     `form:"shuffle" binding="omitempty"`
		Cards      []string `form:"cards" binding="omitempty"`
	}

	DrawCardsFromDeckRequestDTO struct {
		Count uint `form:"count" binding="omitempty"`
	}
)

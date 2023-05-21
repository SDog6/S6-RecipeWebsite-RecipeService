package model

import (
	"time"
)

type RecipePost struct {
	ID           int64
	AuthorID     int64
	Title        string
	Description  string
	Ingredients  string
	Instructions string
	CreatedAt    time.Time
}

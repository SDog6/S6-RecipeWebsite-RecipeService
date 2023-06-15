package model

import (
	"time"
)

type RecipePost struct {
	ID          int64        `json:"id"`
	Author      string     `json:"author"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Picture string     `json:"picture"`
	Ingredients string     `json:"ingredients"`
	Instructions string    `json:"instructions"`
	CreatedAt   *time.Time `json:"created_at"`
}


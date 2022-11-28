package domain

import "time"

type Cake struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title" binding:"required,max=255"`
	Description string     `json:"description" binding:"required"`
	Rating      float32    `json:"rating" binding:"required,gte=1,lte=10"`
	Image       string     `json:"image" binding:"max=255"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

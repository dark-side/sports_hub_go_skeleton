package model

import (
	"time"
)

type Like struct {
	ID           uint `gorm:"primaryKey"`
	Likes        int
	Dislikes     int
	LikeableType string
	LikeableID   uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

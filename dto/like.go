package dto

import "time"

type LikeCreateRequest struct {
	LikeableType string `json:"likeable_type" binding:"required"`
	LikeableID   uint   `json:"likeable_id" binding:"required"`
}

type LikeResponse struct {
	ID           uint      `json:"id"`
	Likes        int       `json:"likes"`
	Dislikes     int       `json:"dislikes"`
	LikeableType string    `json:"likeable_type"`
	LikeableID   uint      `json:"likeable_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

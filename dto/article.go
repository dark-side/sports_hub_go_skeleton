package dto

import "time"

type ArticleCreateRequest struct {
	Title            string `json:"title" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	UserID           uint64 `json:"-"`
}

type ArticleUpdateRequest struct {
	ID               uint   `json:"-"`
	Title            string `json:"title" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	UserID           uint64 `json:"-"`
}

type ArticleResponse struct {
	ID               uint      `json:"id"`
	Title            string    `json:"title"`
	ShortDescription string    `json:"short_description"`
	Description      string    `json:"description"`
	UserID           uint64    `json:"user_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

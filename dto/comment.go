package dto

import "time"

type CommentCreateRequest struct {
	Text      string `json:"text" binding:"required"`
	ArticleID uint   `json:"article_id" binding:"required"`
}

type CommentResponse struct {
	ID        uint      `json:"id"`
	Text      string    `json:"text"`
	UserID    uint      `json:"user_id"`
	ArticleID uint      `json:"article_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

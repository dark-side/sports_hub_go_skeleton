package repository

import (
	"go_be_plgrnd/model"

	"gorm.io/gorm"
)

type CommentRepository interface {
	InsertComment(comment model.Comment) model.Comment
	UpdateComment(comment model.Comment) model.Comment
	DeleteComment(comment model.Comment) error
	FindCommentByID(commentID uint) model.Comment
	FindCommentsByArticleID(articleID uint) []model.Comment
}

type commentConnection struct {
	connection *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentConnection{
		connection: db,
	}
}

func (db *commentConnection) InsertComment(comment model.Comment) model.Comment {
	db.connection.Create(&comment)
	return comment
}

func (db *commentConnection) UpdateComment(comment model.Comment) model.Comment {
	db.connection.Save(&comment)
	return comment
}

func (db *commentConnection) DeleteComment(comment model.Comment) error {
	return db.connection.Delete(&comment).Error
}

func (db *commentConnection) FindCommentByID(commentID uint) model.Comment {
	var comment model.Comment
	db.connection.First(&comment, commentID)
	return comment
}

func (db *commentConnection) FindCommentsByArticleID(articleID uint) []model.Comment {
	var comments []model.Comment
	db.connection.Where("article_id = ?", articleID).Find(&comments)
	return comments
}

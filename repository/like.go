package repository

import (
	"go_be_plgrnd/model"

	"gorm.io/gorm"
)

type LikeRepository interface {
	InsertLike(like model.Like) model.Like
	UpdateLike(like model.Like) model.Like
	DeleteLike(like model.Like) error
	FindLikeByID(likeID uint) model.Like
	FindLikesByLikeable(likeableType string, likeableID uint) []model.Like
}

type likeConnection struct {
	connection *gorm.DB
}

func NewLikeRepository(db *gorm.DB) LikeRepository {
	return &likeConnection{
		connection: db,
	}
}

func (db *likeConnection) InsertLike(like model.Like) model.Like {
	db.connection.Create(&like)
	return like
}

func (db *likeConnection) UpdateLike(like model.Like) model.Like {
	db.connection.Save(&like)
	return like
}

func (db *likeConnection) DeleteLike(like model.Like) error {
	return db.connection.Delete(&like).Error
}

func (db *likeConnection) FindLikeByID(likeID uint) model.Like {
	var like model.Like
	db.connection.First(&like, likeID)
	return like
}

func (db *likeConnection) FindLikesByLikeable(likeableType string, likeableID uint) []model.Like {
	var likes []model.Like
	db.connection.Where("likeable_type = ? AND likeable_id = ?", likeableType, likeableID).Find(&likes)
	return likes
}

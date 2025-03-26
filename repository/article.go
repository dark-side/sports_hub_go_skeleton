package repository

import (
	"go_be_plgrnd/model"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	InsertArticle(article model.Article) model.Article
	UpdateArticle(article model.Article) model.Article
	DeleteArticle(article model.Article) error
	FindArticleByID(articleID uint) model.Article
	FindAllArticles() []model.Article
}

type articleConnection struct {
	connection *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleConnection{
		connection: db,
	}
}

func (db *articleConnection) InsertArticle(article model.Article) model.Article {
	db.connection.Create(&article)
	return article
}

func (db *articleConnection) UpdateArticle(article model.Article) model.Article {
	db.connection.Save(&article)
	return article
}

func (db *articleConnection) DeleteArticle(article model.Article) error {
	return db.connection.Delete(&article).Error
}

func (db *articleConnection) FindArticleByID(articleID uint) model.Article {
	var article model.Article
	db.connection.First(&article, articleID)
	return article
}

func (db *articleConnection) FindAllArticles() []model.Article {
	var articles []model.Article
	db.connection.Find(&articles)
	return articles
}

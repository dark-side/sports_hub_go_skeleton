package service

import (
	"log"

	"go_be_plgrnd/dto"
	"go_be_plgrnd/model"
	"go_be_plgrnd/repository"

	"github.com/mashingan/smapping"
)

type ArticleService interface {
	Create(article dto.ArticleCreateRequest) model.Article
	Update(article dto.ArticleUpdateRequest) model.Article
	Delete(articleID uint) error
	GetArticleByID(articleID uint) model.Article
	GetAllArticles() []model.Article
}

type articleService struct {
	articleRepository repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &articleService{
		articleRepository: repo,
	}
}

func (s *articleService) Create(articleDTO dto.ArticleCreateRequest) model.Article {
	article := model.Article{}
	err := smapping.FillStruct(&article, smapping.MapFields(&articleDTO))
	if err != nil {
		log.Fatalf("Failed to map ArticleCreateRequest: %v", err)
	}
	return s.articleRepository.InsertArticle(article)
}

func (s *articleService) Update(articleDTO dto.ArticleUpdateRequest) model.Article {
	article := model.Article{}
	err := smapping.FillStruct(&article, smapping.MapFields(&articleDTO))
	if err != nil {
		log.Fatalf("Failed to map ArticleUpdateRequest: %v", err)
	}
	article.ID = uint(articleDTO.ID)
	article.UserID = articleDTO.UserID

	return s.articleRepository.UpdateArticle(article)
}

func (s *articleService) Delete(articleID uint) error {
	article := s.articleRepository.FindArticleByID(articleID)
	return s.articleRepository.DeleteArticle(article)
}

func (s *articleService) GetArticleByID(articleID uint) model.Article {
	return s.articleRepository.FindArticleByID(articleID)
}

func (s *articleService) GetAllArticles() []model.Article {
	return s.articleRepository.FindAllArticles()
}

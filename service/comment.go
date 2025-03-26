package service

import (
	"log"

	"go_be_plgrnd/dto"
	"go_be_plgrnd/model"
	"go_be_plgrnd/repository"

	"github.com/mashingan/smapping"
)

type CommentService interface {
	Create(comment dto.CommentCreateRequest) model.Comment
	Update(comment dto.CommentCreateRequest) model.Comment
	Delete(commentID uint) error
	GetCommentByID(commentID uint) model.Comment
	GetCommentsByArticleID(articleID uint) []model.Comment
}

type commentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) CommentService {
	return &commentService{
		commentRepository: repo,
	}
}

func (s *commentService) Create(commentDTO dto.CommentCreateRequest) model.Comment {
	comment := model.Comment{}
	err := smapping.FillStruct(&comment, smapping.MapFields(&commentDTO))
	if err != nil {
		log.Fatalf("Failed to map CommentCreateRequest: %v", err)
	}
	return s.commentRepository.InsertComment(comment)
}

func (s *commentService) Update(commentDTO dto.CommentCreateRequest) model.Comment {
	comment := model.Comment{}
	err := smapping.FillStruct(&comment, smapping.MapFields(&commentDTO))
	if err != nil {
		log.Fatalf("Failed to map CommentCreateRequest: %v", err)
	}
	return s.commentRepository.UpdateComment(comment)
}

func (s *commentService) Delete(commentID uint) error {
	comment := s.commentRepository.FindCommentByID(commentID)
	return s.commentRepository.DeleteComment(comment)
}

func (s *commentService) GetCommentByID(commentID uint) model.Comment {
	return s.commentRepository.FindCommentByID(commentID)
}

func (s *commentService) GetCommentsByArticleID(articleID uint) []model.Comment {
	return s.commentRepository.FindCommentsByArticleID(articleID)
}

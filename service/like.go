package service

import (
	"log"

	"go_be_plgrnd/dto"
	"go_be_plgrnd/model"
	"go_be_plgrnd/repository"

	"github.com/mashingan/smapping"
)

type LikeService interface {
	Create(like dto.LikeCreateRequest) model.Like
	Update(like dto.LikeCreateRequest) model.Like
	Delete(likeID uint) error
	GetLikeByID(likeID uint) model.Like
	GetLikesByLikeable(likeableType string, likeableID uint) []model.Like
}

type likeService struct {
	likeRepository repository.LikeRepository
}

func NewLikeService(repo repository.LikeRepository) LikeService {
	return &likeService{
		likeRepository: repo,
	}
}

func (s *likeService) Create(likeDTO dto.LikeCreateRequest) model.Like {
	like := model.Like{}
	err := smapping.FillStruct(&like, smapping.MapFields(&likeDTO))
	if err != nil {
		log.Fatalf("Failed to map LikeCreateRequest: %v", err)
	}
	return s.likeRepository.InsertLike(like)
}

func (s *likeService) Update(likeDTO dto.LikeCreateRequest) model.Like {
	like := model.Like{}
	err := smapping.FillStruct(&like, smapping.MapFields(&likeDTO))
	if err != nil {
		log.Fatalf("Failed to map LikeCreateRequest: %v", err)
	}
	return s.likeRepository.UpdateLike(like)
}

func (s *likeService) Delete(likeID uint) error {
	like := s.likeRepository.FindLikeByID(likeID)
	return s.likeRepository.DeleteLike(like)
}

func (s *likeService) GetLikeByID(likeID uint) model.Like {
	return s.likeRepository.FindLikeByID(likeID)
}

func (s *likeService) GetLikesByLikeable(likeableType string, likeableID uint) []model.Like {
	return s.likeRepository.FindLikesByLikeable(likeableType, likeableID)
}

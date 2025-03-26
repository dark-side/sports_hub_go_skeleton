package service

import (
	"go_be_plgrnd/model"
	"go_be_plgrnd/repository"
)

type ActiveStorageBlobService interface {
	Create(blob model.ActiveStorageBlob) model.ActiveStorageBlob
	Update(blob model.ActiveStorageBlob) model.ActiveStorageBlob
	Delete(blobID uint) error
	GetByID(blobID uint) model.ActiveStorageBlob
	GetByKey(key string) model.ActiveStorageBlob
}

type activeStorageBlobService struct {
	repo repository.ActiveStorageBlobRepository
}

func NewActiveStorageBlobService(repo repository.ActiveStorageBlobRepository) ActiveStorageBlobService {
	return &activeStorageBlobService{
		repo: repo,
	}
}

func (s *activeStorageBlobService) Create(blob model.ActiveStorageBlob) model.ActiveStorageBlob {
	return s.repo.InsertBlob(blob)
}

func (s *activeStorageBlobService) Update(blob model.ActiveStorageBlob) model.ActiveStorageBlob {
	return s.repo.UpdateBlob(blob)
}

func (s *activeStorageBlobService) Delete(blobID uint) error {
	blob := s.repo.FindBlobByID(blobID)
	return s.repo.DeleteBlob(blob)
}

func (s *activeStorageBlobService) GetByID(blobID uint) model.ActiveStorageBlob {
	return s.repo.FindBlobByID(blobID)
}

func (s *activeStorageBlobService) GetByKey(key string) model.ActiveStorageBlob {
	return s.repo.FindBlobByKey(key)
}

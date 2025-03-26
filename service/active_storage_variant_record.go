package service

import (
	"go_be_plgrnd/model"
	"go_be_plgrnd/repository"
)

type ActiveStorageVariantRecordService interface {
	Create(variant model.ActiveStorageVariantRecord) model.ActiveStorageVariantRecord
	Update(variant model.ActiveStorageVariantRecord) model.ActiveStorageVariantRecord
	Delete(variantID uint) error
	GetByID(variantID uint) model.ActiveStorageVariantRecord
	GetByBlobID(blobID uint) []model.ActiveStorageVariantRecord
}

type activeStorageVariantRecordService struct {
	repo repository.ActiveStorageVariantRecordRepository
}

func NewActiveStorageVariantRecordService(repo repository.ActiveStorageVariantRecordRepository) ActiveStorageVariantRecordService {
	return &activeStorageVariantRecordService{
		repo: repo,
	}
}

func (s *activeStorageVariantRecordService) Create(variant model.ActiveStorageVariantRecord) model.ActiveStorageVariantRecord {
	return s.repo.InsertVariant(variant)
}

func (s *activeStorageVariantRecordService) Update(variant model.ActiveStorageVariantRecord) model.ActiveStorageVariantRecord {
	return s.repo.UpdateVariant(variant)
}

func (s *activeStorageVariantRecordService) Delete(variantID uint) error {
	variant := s.repo.FindVariantByID(variantID)
	return s.repo.DeleteVariant(variant)
}

func (s *activeStorageVariantRecordService) GetByID(variantID uint) model.ActiveStorageVariantRecord {
	return s.repo.FindVariantByID(variantID)
}

func (s *activeStorageVariantRecordService) GetByBlobID(blobID uint) []model.ActiveStorageVariantRecord {
	return s.repo.FindVariantsByBlobID(blobID)
}

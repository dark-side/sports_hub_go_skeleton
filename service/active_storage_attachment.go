package service

import (
	"go_be_plgrnd/model"
	"go_be_plgrnd/repository"
)

type ActiveStorageAttachmentService interface {
	Create(attachment model.ActiveStorageAttachment) model.ActiveStorageAttachment
	Update(attachment model.ActiveStorageAttachment) model.ActiveStorageAttachment
	Delete(attachmentID uint) error
	GetByID(attachmentID uint) model.ActiveStorageAttachment
	GetByRecord(recordType string, recordID uint) []model.ActiveStorageAttachment
}

type activeStorageAttachmentService struct {
	repo repository.ActiveStorageAttachmentRepository
}

func NewActiveStorageAttachmentService(repo repository.ActiveStorageAttachmentRepository) ActiveStorageAttachmentService {
	return &activeStorageAttachmentService{
		repo: repo,
	}
}

func (s *activeStorageAttachmentService) Create(attachment model.ActiveStorageAttachment) model.ActiveStorageAttachment {
	return s.repo.InsertAttachment(attachment)
}

func (s *activeStorageAttachmentService) Update(attachment model.ActiveStorageAttachment) model.ActiveStorageAttachment {
	return s.repo.UpdateAttachment(attachment)
}

func (s *activeStorageAttachmentService) Delete(attachmentID uint) error {
	attachment := s.repo.FindAttachmentByID(attachmentID)
	return s.repo.DeleteAttachment(attachment)
}

func (s *activeStorageAttachmentService) GetByID(attachmentID uint) model.ActiveStorageAttachment {
	return s.repo.FindAttachmentByID(attachmentID)
}

func (s *activeStorageAttachmentService) GetByRecord(recordType string, recordID uint) []model.ActiveStorageAttachment {
	return s.repo.FindAttachmentsByRecord(recordType, recordID)
}

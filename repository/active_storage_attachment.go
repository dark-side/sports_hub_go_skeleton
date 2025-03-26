package repository

import (
	"go_be_plgrnd/model"

	"gorm.io/gorm"
)

type ActiveStorageAttachmentRepository interface {
	InsertAttachment(attachment model.ActiveStorageAttachment) model.ActiveStorageAttachment
	UpdateAttachment(attachment model.ActiveStorageAttachment) model.ActiveStorageAttachment
	DeleteAttachment(attachment model.ActiveStorageAttachment) error
	FindAttachmentByID(attachmentID uint) model.ActiveStorageAttachment
	FindAttachmentsByRecord(recordType string, recordID uint) []model.ActiveStorageAttachment
}

type attachmentConnection struct {
	connection *gorm.DB
}

func NewActiveStorageAttachmentRepository(db *gorm.DB) ActiveStorageAttachmentRepository {
	return &attachmentConnection{
		connection: db,
	}
}

func (db *attachmentConnection) InsertAttachment(attachment model.ActiveStorageAttachment) model.ActiveStorageAttachment {
	db.connection.Create(&attachment)
	return attachment
}

func (db *attachmentConnection) UpdateAttachment(attachment model.ActiveStorageAttachment) model.ActiveStorageAttachment {
	db.connection.Save(&attachment)
	return attachment
}

func (db *attachmentConnection) DeleteAttachment(attachment model.ActiveStorageAttachment) error {
	return db.connection.Delete(&attachment).Error
}

func (db *attachmentConnection) FindAttachmentByID(attachmentID uint) model.ActiveStorageAttachment {
	var attachment model.ActiveStorageAttachment
	db.connection.First(&attachment, attachmentID)
	return attachment
}

func (db *attachmentConnection) FindAttachmentsByRecord(recordType string, recordID uint) []model.ActiveStorageAttachment {
	var attachments []model.ActiveStorageAttachment
	db.connection.Where("record_type = ? AND record_id = ?", recordType, recordID).Find(&attachments)
	return attachments
}

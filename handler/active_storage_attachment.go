package handler

import (
	"net/http"
	"strconv"

	"go_be_plgrnd/helper"
	"go_be_plgrnd/model"
	"go_be_plgrnd/service"

	"github.com/gin-gonic/gin"
)

type ActiveStorageAttachmentHandler interface {
	CreateAttachment(ctx *gin.Context)
	UpdateAttachment(ctx *gin.Context)
	DeleteAttachment(ctx *gin.Context)
	GetAttachmentByID(ctx *gin.Context)
	GetAttachmentsByRecord(ctx *gin.Context)
}

type activeStorageAttachmentHandler struct {
	attachmentService service.ActiveStorageAttachmentService
	jwtService        service.JWTService
}

func NewActiveStorageAttachmentHandler(attachmentService service.ActiveStorageAttachmentService, jwtService service.JWTService) ActiveStorageAttachmentHandler {
	return &activeStorageAttachmentHandler{
		attachmentService: attachmentService,
		jwtService:        jwtService,
	}
}

func (c *activeStorageAttachmentHandler) CreateAttachment(ctx *gin.Context) {
	var attachment model.ActiveStorageAttachment
	if err := ctx.ShouldBindJSON(&attachment); err != nil {
		res := helper.BuildErrorResponse("Failed to bind request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	createdAttachment := c.attachmentService.Create(attachment)
	res := helper.BuildResponse(true, "Attachment created successfully", createdAttachment)
	ctx.JSON(http.StatusOK, res)
}

func (c *activeStorageAttachmentHandler) UpdateAttachment(ctx *gin.Context) {
	var attachment model.ActiveStorageAttachment
	if err := ctx.ShouldBindJSON(&attachment); err != nil {
		res := helper.BuildErrorResponse("Failed to bind request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	updatedAttachment := c.attachmentService.Update(attachment)
	res := helper.BuildResponse(true, "Attachment updated successfully", updatedAttachment)
	ctx.JSON(http.StatusOK, res)
}

func (c *activeStorageAttachmentHandler) DeleteAttachment(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid attachment ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	err = c.attachmentService.Delete(uint(id))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to delete attachment", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	res := helper.BuildResponse(true, "Attachment deleted successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *activeStorageAttachmentHandler) GetAttachmentByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid attachment ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	attachment := c.attachmentService.GetByID(uint(id))
	res := helper.BuildResponse(true, "OK", attachment)
	ctx.JSON(http.StatusOK, res)
}

func (c *activeStorageAttachmentHandler) GetAttachmentsByRecord(ctx *gin.Context) {
	recordType := ctx.Param("type")
	idParam := ctx.Param("id")
	recordID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid record ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	attachments := c.attachmentService.GetByRecord(recordType, uint(recordID))
	res := helper.BuildResponse(true, "OK", attachments)
	ctx.JSON(http.StatusOK, res)
}

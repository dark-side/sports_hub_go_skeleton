package handler

import (
	"net/http"
	"strconv"

	"go_be_plgrnd/helper"
	"go_be_plgrnd/model"
	"go_be_plgrnd/service"

	"github.com/gin-gonic/gin"
)

type ActiveStorageVariantRecordHandler interface {
	CreateVariantRecord(ctx *gin.Context)
	UpdateVariantRecord(ctx *gin.Context)
	DeleteVariantRecord(ctx *gin.Context)
	GetVariantRecordByID(ctx *gin.Context)
	GetVariantRecordsByBlobID(ctx *gin.Context)
}

type activeStorageVariantRecordHandler struct {
	variantService service.ActiveStorageVariantRecordService
	jwtService     service.JWTService
}

func NewActiveStorageVariantRecordHandler(variantService service.ActiveStorageVariantRecordService, jwtService service.JWTService) ActiveStorageVariantRecordHandler {
	return &activeStorageVariantRecordHandler{
		variantService: variantService,
		jwtService:     jwtService,
	}
}

func (c *activeStorageVariantRecordHandler) CreateVariantRecord(ctx *gin.Context) {
	var variant model.ActiveStorageVariantRecord
	if err := ctx.ShouldBindJSON(&variant); err != nil {
		res := helper.BuildErrorResponse("Failed to bind request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	createdVariant := c.variantService.Create(variant)
	res := helper.BuildResponse(true, "Variant record created successfully", createdVariant)
	ctx.JSON(http.StatusOK, res)
}

func (c *activeStorageVariantRecordHandler) UpdateVariantRecord(ctx *gin.Context) {
	var variant model.ActiveStorageVariantRecord
	if err := ctx.ShouldBindJSON(&variant); err != nil {
		res := helper.BuildErrorResponse("Failed to bind request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	updatedVariant := c.variantService.Update(variant)
	res := helper.BuildResponse(true, "Variant record updated successfully", updatedVariant)
	ctx.JSON(http.StatusOK, res)
}

func (c *activeStorageVariantRecordHandler) DeleteVariantRecord(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid variant record ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	err = c.variantService.Delete(uint(id))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to delete variant record", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	res := helper.BuildResponse(true, "Variant record deleted successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *activeStorageVariantRecordHandler) GetVariantRecordByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid variant record ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	variant := c.variantService.GetByID(uint(id))
	res := helper.BuildResponse(true, "OK", variant)
	ctx.JSON(http.StatusOK, res)
}

func (c *activeStorageVariantRecordHandler) GetVariantRecordsByBlobID(ctx *gin.Context) {
	blobIdParam := ctx.Param("blobId")
	blobId, err := strconv.ParseUint(blobIdParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid blob ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	variants := c.variantService.GetByBlobID(uint(blobId))
	res := helper.BuildResponse(true, "OK", variants)
	ctx.JSON(http.StatusOK, res)
}

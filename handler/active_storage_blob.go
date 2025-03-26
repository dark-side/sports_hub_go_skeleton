package handler

import (
	"net/http"
	"strconv"

	"go_be_plgrnd/helper"
	"go_be_plgrnd/model"
	"go_be_plgrnd/service"

	"github.com/gin-gonic/gin"
)

type ActiveStorageBlobHandler interface {
	CreateBlob(ctx *gin.Context)
	UpdateBlob(ctx *gin.Context)
	DeleteBlob(ctx *gin.Context)
	GetBlobByID(ctx *gin.Context)
	GetBlobByKey(ctx *gin.Context)
}

type activeStorageBlobHandler struct {
	blobService service.ActiveStorageBlobService
	jwtService  service.JWTService
}

func NewActiveStorageBlobHandler(blobService service.ActiveStorageBlobService, jwtService service.JWTService) ActiveStorageBlobHandler {
	return &activeStorageBlobHandler{
		blobService: blobService,
		jwtService:  jwtService,
	}
}

func (c *activeStorageBlobHandler) CreateBlob(ctx *gin.Context) {
	var blob model.ActiveStorageBlob
	if err := ctx.ShouldBindJSON(&blob); err != nil {
		res := helper.BuildErrorResponse("Failed to bind request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	createdBlob := c.blobService.Create(blob)
	res := helper.BuildResponse(true, "Blob created successfully", createdBlob)
	ctx.JSON(http.StatusOK, res)
}

func (c *activeStorageBlobHandler) UpdateBlob(ctx *gin.Context) {
	var blob model.ActiveStorageBlob
	if err := ctx.ShouldBindJSON(&blob); err != nil {
		res := helper.BuildErrorResponse("Failed to bind request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	updatedBlob := c.blobService.Update(blob)
	res := helper.BuildResponse(true, "Blob updated successfully", updatedBlob)
	ctx.JSON(http.StatusOK, res)
}

func (c *activeStorageBlobHandler) DeleteBlob(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid blob ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	err = c.blobService.Delete(uint(id))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to delete blob", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	res := helper.BuildResponse(true, "Blob deleted successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *activeStorageBlobHandler) GetBlobByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid blob ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	blob := c.blobService.GetByID(uint(id))
	res := helper.BuildResponse(true, "OK", blob)
	ctx.JSON(http.StatusOK, res)
}

func (c *activeStorageBlobHandler) GetBlobByKey(ctx *gin.Context) {
	key := ctx.Param("key")
	blob := c.blobService.GetByKey(key)
	res := helper.BuildResponse(true, "OK", blob)
	ctx.JSON(http.StatusOK, res)
}

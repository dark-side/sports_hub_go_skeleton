package handler

import (
	"net/http"
	"strconv"

	"go_be_plgrnd/dto"
	"go_be_plgrnd/helper"
	"go_be_plgrnd/service"

	"github.com/gin-gonic/gin"
)

type LikeHandler interface {
	CreateLike(ctx *gin.Context)
	UpdateLike(ctx *gin.Context)
	DeleteLike(ctx *gin.Context)
	GetLikeByID(ctx *gin.Context)
	GetLikesByLikeable(ctx *gin.Context)
}

type likeHandler struct {
	likeService service.LikeService
	jwtService  service.JWTService
}

func NewLikeHandler(likeService service.LikeService, jwtService service.JWTService) LikeHandler {
	return &likeHandler{
		likeService: likeService,
		jwtService:  jwtService,
	}
}

func (c *likeHandler) CreateLike(ctx *gin.Context) {
	var likeDTO dto.LikeCreateRequest
	if err := ctx.ShouldBindJSON(&likeDTO); err != nil {
		res := helper.BuildErrorResponse("Failed to bind request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	like := c.likeService.Create(likeDTO)
	res := helper.BuildResponse(true, "Like created successfully", like)
	ctx.JSON(http.StatusOK, res)
}

func (c *likeHandler) UpdateLike(ctx *gin.Context) {
	var likeDTO dto.LikeCreateRequest
	if err := ctx.ShouldBindJSON(&likeDTO); err != nil {
		res := helper.BuildErrorResponse("Failed to bind request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	idParam := ctx.Param("id")
	_, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid like ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	like := c.likeService.Update(likeDTO)
	res := helper.BuildResponse(true, "Like updated successfully", like)
	ctx.JSON(http.StatusOK, res)
}

func (c *likeHandler) DeleteLike(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid like ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	err = c.likeService.Delete(uint(id))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to delete like", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	res := helper.BuildResponse(true, "Like deleted successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *likeHandler) GetLikeByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid like ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	like := c.likeService.GetLikeByID(uint(id))
	res := helper.BuildResponse(true, "OK", like)
	ctx.JSON(http.StatusOK, res)
}

func (c *likeHandler) GetLikesByLikeable(ctx *gin.Context) {
	likeableType := ctx.Param("type")
	idParam := ctx.Param("id")
	likeableID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid likeable ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	likes := c.likeService.GetLikesByLikeable(likeableType, uint(likeableID))
	res := helper.BuildResponse(true, "OK", likes)
	ctx.JSON(http.StatusOK, res)
}

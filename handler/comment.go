package handler

import (
	"net/http"
	"strconv"

	"go_be_plgrnd/dto"
	"go_be_plgrnd/helper"
	"go_be_plgrnd/service"

	"github.com/gin-gonic/gin"
)

type CommentHandler interface {
	CreateComment(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
	GetCommentByID(ctx *gin.Context)
	GetCommentsByArticleID(ctx *gin.Context)
}

type commentHandler struct {
	commentService service.CommentService
	jwtService     service.JWTService
}

func NewCommentHandler(commentService service.CommentService, jwtService service.JWTService) CommentHandler {
	return &commentHandler{
		commentService: commentService,
		jwtService:     jwtService,
	}
}

func (c *commentHandler) CreateComment(ctx *gin.Context) {
	var commentDTO dto.CommentCreateRequest
	if err := ctx.ShouldBindJSON(&commentDTO); err != nil {
		res := helper.BuildErrorResponse("Failed to bind request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	comment := c.commentService.Create(commentDTO)
	res := helper.BuildResponse(true, "Comment created successfully", comment)
	ctx.JSON(http.StatusOK, res)
}

func (c *commentHandler) UpdateComment(ctx *gin.Context) {
	var commentDTO dto.CommentCreateRequest
	if err := ctx.ShouldBindJSON(&commentDTO); err != nil {
		res := helper.BuildErrorResponse("Failed to bind request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	idParam := ctx.Param("id")
	_, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid comment ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	comment := c.commentService.Update(commentDTO)
	res := helper.BuildResponse(true, "Comment updated successfully", comment)
	ctx.JSON(http.StatusOK, res)
}

func (c *commentHandler) DeleteComment(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid comment ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	err = c.commentService.Delete(uint(id))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to delete comment", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	res := helper.BuildResponse(true, "Comment deleted successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *commentHandler) GetCommentByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid comment ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	comment := c.commentService.GetCommentByID(uint(id))
	res := helper.BuildResponse(true, "OK", comment)
	ctx.JSON(http.StatusOK, res)
}

func (c *commentHandler) GetCommentsByArticleID(ctx *gin.Context) {
	articleIdParam := ctx.Param("articleId")
	articleId, err := strconv.ParseUint(articleIdParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid article ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	comments := c.commentService.GetCommentsByArticleID(uint(articleId))
	res := helper.BuildResponse(true, "OK", comments)
	ctx.JSON(http.StatusOK, res)
}

package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"go_be_plgrnd/dto"
	"go_be_plgrnd/helper"
	"go_be_plgrnd/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ArticleHandler interface {
	CreateArticle(ctx *gin.Context)
	UpdateArticle(ctx *gin.Context)
	DeleteArticle(ctx *gin.Context)
	GetArticleByID(ctx *gin.Context)
	GetAllArticles(ctx *gin.Context)
}

type articleHandler struct {
	articleService service.ArticleService
	jwtService     service.JWTService
}

func NewArticleHandler(articleService service.ArticleService, jwtService service.JWTService) ArticleHandler {
	return &articleHandler{
		articleService: articleService,
		jwtService:     jwtService,
	}
}

func (c *articleHandler) CreateArticle(ctx *gin.Context) {
	var articleDTO dto.ArticleCreateRequest
	if err := ctx.ShouldBindJSON(&articleDTO); err != nil {
		res := helper.BuildErrorResponse("Failed to bind request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		res := helper.BuildErrorResponse("Token validation failed", errToken.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userID, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to parse user_id from token", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	articleDTO.UserID = userID

	article := c.articleService.Create(articleDTO)

	res := helper.BuildResponse(true, "Article created successfully", article)
	ctx.JSON(http.StatusOK, res)
}

func (c *articleHandler) UpdateArticle(ctx *gin.Context) {
	var articleDTO dto.ArticleUpdateRequest
	if err := ctx.ShouldBindJSON(&articleDTO); err != nil {
		res := helper.BuildErrorResponse("Failed to bind request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid article ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	articleDTO.ID = uint(id)

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		res := helper.BuildErrorResponse("Token validation failed", errToken.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to parse user_id from token", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	articleDTO.UserID = userID

	article := c.articleService.Update(articleDTO)
	res := helper.BuildResponse(true, "Article updated successfully", article)
	ctx.JSON(http.StatusOK, res)
}

func (c *articleHandler) DeleteArticle(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid article ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	err = c.articleService.Delete(uint(id))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to delete article", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	res := helper.BuildResponse(true, "Article deleted successfully", nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *articleHandler) GetArticleByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Invalid article ID", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	article := c.articleService.GetArticleByID(uint(id))
	res := helper.BuildResponse(true, "OK", article)
	ctx.JSON(http.StatusOK, res)
}

func (c *articleHandler) GetAllArticles(ctx *gin.Context) {
	articles := c.articleService.GetAllArticles()
	res := helper.BuildResponse(true, "OK", articles)
	ctx.JSON(http.StatusOK, res)
}

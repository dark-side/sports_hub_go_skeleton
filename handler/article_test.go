// article_handler_test.go
package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"go_be_plgrnd/dto"
	"go_be_plgrnd/model"
	"go_be_plgrnd/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Test_suit(t *testing.T) {
	t.Run("TestCreateArticle", TestCreateArticle)
	t.Run("TestUpdateArticle", TestUpdateArticle)
	t.Run("TestDeleteArticle", TestDeleteArticle)
	t.Run("TestGetArticleByID", TestGetArticleByID)
	t.Run("TestGetAllArticles", TestGetAllArticles)
}

func NewArticleHandlerT(articleService service.ArticleService, jwtService service.JWTService) ArticleHandler {
	return NewArticleHandler(articleService, jwtService)
}

type FakeArticleService struct{}

func (f *FakeArticleService) Create(articleDTO dto.ArticleCreateRequest) model.Article {
	return model.Article{
		ID:               1,
		Title:            articleDTO.Title,
		ShortDescription: articleDTO.ShortDescription,
		Description:      articleDTO.Description,
		UserID:           articleDTO.UserID,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

func (f *FakeArticleService) Update(articleDTO dto.ArticleUpdateRequest) model.Article {
	return model.Article{
		ID:               uint(articleDTO.ID),
		Title:            articleDTO.Title,
		ShortDescription: articleDTO.ShortDescription,
		Description:      articleDTO.Description,
		UserID:           articleDTO.UserID,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

func (f *FakeArticleService) Delete(articleID uint) error {
	return nil
}

func (f *FakeArticleService) GetArticleByID(articleID uint) model.Article {
	return model.Article{
		ID:               articleID,
		Title:            "Test Article",
		ShortDescription: "Test Short Description",
		Description:      "Test Description",
		UserID:           1,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

func (f *FakeArticleService) GetAllArticles() []model.Article {
	return []model.Article{
		{
			ID:               1,
			Title:            "Test Article",
			ShortDescription: "Test Short Description",
			Description:      "Test Description",
			UserID:           1,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
	}
}

type FakeJWTService struct{}

func (f *FakeJWTService) GenerateToken(userID string) string {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func (f *FakeJWTService) ValidateToken(tokenStr string) (*jwt.Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": "1",
	})
	return token, nil
}

func setupRouter(handler ArticleHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/articles", handler.CreateArticle)
	r.PUT("/api/articles/:id", handler.UpdateArticle)
	r.DELETE("/api/articles/:id", handler.DeleteArticle)
	r.GET("/api/articles/:id", handler.GetArticleByID)
	r.GET("/api/articles", handler.GetAllArticles)
	return r
}

func TestCreateArticle(t *testing.T) {
	fakeArticleService := &FakeArticleService{}
	fakeJWTService := &FakeJWTService{}
	handler := NewArticleHandlerT(fakeArticleService, fakeJWTService)
	router := setupRouter(handler)

	articlePayload := dto.ArticleCreateRequest{
		Title:            "New Article",
		ShortDescription: "Short Desc",
		Description:      "Full description",
	}
	body, _ := json.Marshal(articlePayload)
	req, _ := http.NewRequest("POST", "/api/articles", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer fakeToken")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(resp.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	successVal, exists := response["success"]
	if !exists {
		t.Errorf("Expected key 'success' in response, got nil")
	} else if successBool, ok := successVal.(bool); !ok || successBool != true {
		t.Errorf("Expected success=true, got %v", successVal)
	}

	data, exists := response["data"].(map[string]interface{})
	if !exists {
		t.Errorf("Expected data to be a map, got %v", response["data"])
	} else {
		userIDVal, exists := data["user_id"]
		if !exists {
			t.Errorf("Expected user_id in data, got nil")
		} else if userIDFloat, ok := userIDVal.(float64); !ok || uint64(userIDFloat) != 1 {
			t.Errorf("Expected user_id to be 1, got %v", userIDVal)
		}
	}
}

func TestUpdateArticle(t *testing.T) {
	fakeArticleService := &FakeArticleService{}
	fakeJWTService := &FakeJWTService{}
	handler := NewArticleHandlerT(fakeArticleService, fakeJWTService)
	router := setupRouter(handler)

	articlePayload := dto.ArticleUpdateRequest{
		Title:            "Updated Article",
		ShortDescription: "Updated Short Desc",
		Description:      "Updated full description",
	}
	body, _ := json.Marshal(articlePayload)
	req, _ := http.NewRequest("PUT", "/api/articles/2", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer fakeToken")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(resp.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}
	successVal, exists := response["success"]
	if !exists || successVal != true {
		t.Errorf("Expected success=true, got %v", response["success"])
	}

	data, exists := response["data"].(map[string]interface{})
	if !exists {
		t.Errorf("Expected data to be a map, got %T", response["data"])
	} else {
		id, _ := strconv.Atoi("2")
		if int(data["id"].(float64)) != id {
			t.Errorf("Expected id %d, got %v", id, data["id"])
		}
		if userIDVal, ok := data["user_id"].(float64); !ok || uint64(userIDVal) != 1 {
			t.Errorf("Expected user_id to be 1, got %v", data["user_id"])
		}
	}
}

func TestDeleteArticle(t *testing.T) {
	fakeArticleService := &FakeArticleService{}
	fakeJWTService := &FakeJWTService{}
	handler := NewArticleHandlerT(fakeArticleService, fakeJWTService)
	router := setupRouter(handler)

	req, _ := http.NewRequest("DELETE", "/api/articles/3", nil)
	req.Header.Set("Authorization", "Bearer fakeToken")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.Code)
	}
}

func TestGetArticleByID(t *testing.T) {
	fakeArticleService := &FakeArticleService{}
	fakeJWTService := &FakeJWTService{}
	handler := NewArticleHandlerT(fakeArticleService, fakeJWTService)
	router := setupRouter(handler)

	req, _ := http.NewRequest("GET", "/api/articles/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(resp.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}
	if response["success"] == nil {
		t.Errorf("Expected success=true, got nil")
	} else if success, ok := response["success"].(bool); !ok || success != true {
		t.Errorf("Expected success=true, got %v", response["success"])
	}
}

func TestGetAllArticles(t *testing.T) {
	fakeArticleService := &FakeArticleService{}
	fakeJWTService := &FakeJWTService{}
	handler := NewArticleHandlerT(fakeArticleService, fakeJWTService)
	router := setupRouter(handler)

	req, _ := http.NewRequest("GET", "/api/articles", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(resp.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}
	if response["success"] == nil {
		t.Errorf("Expected success=true, got nil")
	} else if success, ok := response["success"].(bool); !ok || success != true {
		t.Errorf("Expected success=true, got %v", response["success"])
	}

	data, exists := response["data"].([]interface{})
	if !exists {
		t.Errorf("Expected data to be a slice, got %T", response["data"])
	} else if len(data) == 0 {
		t.Errorf("Expected at least one article, got 0")
	}
}

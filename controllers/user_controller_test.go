package controllers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// GetAllUsersHandler のユニットテスト
func TestGetAllUsersHandler(t *testing.T) {
	// Echoのインスタンスを作成
	e := echo.New()

	// モックサービスのセットアップ
	mockService := new(services.MockUserService)

	// モックの期待値設定
	mockService.On("GetAllUsers").Return([]models.User{
		{ID: "1", Username: "JohnDoe", Email: "john@example.com"},
		{ID: "2", Username: "JaneDoe", Email: "jane@example.com"},
	}, nil)

	// テスト用のコントローラーを作成
	controller := NewRealUserController(mockService)

	// `/users`ルートのテスト
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// GetAllUsersHandlerを実行し、レスポンスの確認
	if assert.NoError(t, controller.GetAllUsersHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "JohnDoe")
		assert.Contains(t, rec.Body.String(), "JaneDoe")
	}

	// モックが期待通りに呼ばれたかを検証
	mockService.AssertExpectations(t)
}

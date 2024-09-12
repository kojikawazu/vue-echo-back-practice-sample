package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TestUserHandler のユニットテスト
func TestUserHandler(t *testing.T) {
	// Echoのインスタンスを作成
	e := echo.New()

	// `/users`ルートのテスト
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// GetUsersを実行し、レスポンスの確認
	if assert.NoError(t, GetUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "username")
	}
}

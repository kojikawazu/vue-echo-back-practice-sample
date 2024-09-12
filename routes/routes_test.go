package routes

import (
	"backend/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	// Echoのインスタンスを作成
	e := echo.New()

	// ルーティングの初期化
	InitRoutes(e)

	// `/`ルートのテスト
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// TestHandlerを実行し、レスポンスの確認
	if assert.NoError(t, controllers.TestHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello, World!", rec.Body.String())
	}
}

func TestUsersRoute(t *testing.T) {
	// Echoのインスタンスを作成
	e := echo.New()

	// ルーティングの初期化
	InitRoutes(e)

	// `/users`ルートのテスト
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// GetUsersを実行し、レスポンスの確認
	if assert.NoError(t, controllers.GetUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "user") // レスポンスが "user" を含むことを確認
	}
}

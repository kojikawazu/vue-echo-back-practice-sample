package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TestTodoHandler のユニットテスト
func TestTodoHandler(t *testing.T) {
	// Echoのインスタンスを作成
	e := echo.New()

	// `/todos`ルートのテスト
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// GetTodosを実行し、レスポンスの確認
	if assert.NoError(t, GetTodos(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "description")
	}
}

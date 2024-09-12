package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TestTestHandler は TestHandler のユニットテスト
func TestTestHandler(t *testing.T) {
	// Echoインスタンスを作成
	e := echo.New()

	// テスト用リクエストとレスポンスを作成
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	// コンテキストを作成
	c := e.NewContext(req, rec)

	// ハンドラを実行して結果を確認
	if assert.NoError(t, TestHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello, World!", rec.Body.String())
	}
}

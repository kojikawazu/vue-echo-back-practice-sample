package routes

import (
	"backend/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRoutes(t *testing.T) {
	e := echo.New()

	// モックコントローラーの作成
	mockController := new(controllers.MockTodoController)

	// 期待値を設定
	mockController.On("GetTodosHandler", mock.Anything).Return(nil)

	// ルートをセットアップ
	SetupRoutes(e, mockController)

	// テスト用のリクエストとレスポンスの準備
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec := httptest.NewRecorder()

	// Echoのコンテキストを生成してリクエストを実行
	e.ServeHTTP(rec, req)

	// ステータスコードの検証
	assert.Equal(t, http.StatusOK, rec.Code)

	// モックが期待通りに呼ばれたかを検証
	mockController.AssertExpectations(t)
}

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
	mockTodoController := new(controllers.MockTodoController)
	mockTestController := new(controllers.MockTestController)

	// 期待値を設定
	mockTodoController.On("GetTodosHandler", mock.Anything).Return(nil)
	mockTestController.On("TestHandler", mock.Anything).Return(nil)

	// ルートをセットアップ
	SetupRoutes(e, mockTodoController, mockTestController)

	// ------------------------------------------------------------------

	// テスト用のリクエストとレスポンスの準備
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()

	// Echoのコンテキストを生成してリクエストを実行
	e.ServeHTTP(rec, req)

	// ステータスコードの検証
	assert.Equal(t, http.StatusOK, rec.Code)

	// モックが期待通りに呼ばれたかを検証
	mockTestController.AssertExpectations(t)

	// ------------------------------------------------------------------

	// テスト用のリクエストとレスポンスの準備
	req = httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec = httptest.NewRecorder()

	// Echoのコンテキストを生成してリクエストを実行
	e.ServeHTTP(rec, req)

	// ステータスコードの検証
	assert.Equal(t, http.StatusOK, rec.Code)

	// モックが期待通りに呼ばれたかを検証
	mockTodoController.AssertExpectations(t)

	// ------------------------------------------------------------------
}

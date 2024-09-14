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
	mockTestController := new(controllers.MockTestController)
	mockTodoController := new(controllers.MockTodoController)
	mockUserController := new(controllers.MockUserController)

	// 期待値を設定
	mockTestController.On("TestHandler", mock.Anything).Return(nil)
	mockTodoController.On("GetTodosHandler", mock.Anything).Return(nil)
	mockUserController.On("GetAllUsersHandler", mock.Anything).Return(nil)

	// ルートをセットアップ
	SetupRoutes(e, mockTodoController, mockTestController, mockUserController)

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

	// テスト用のリクエストとレスポンスの準備
	req = httptest.NewRequest(http.MethodGet, "/users", nil)
	rec = httptest.NewRecorder()

	// Echoのコンテキストを生成してリクエストを実行
	e.ServeHTTP(rec, req)

	// ステータスコードの検証
	assert.Equal(t, http.StatusOK, rec.Code)

	// モックが期待通りに呼ばれたかを検証
	mockUserController.AssertExpectations(t)

	// ------------------------------------------------------------------
}

package controllers

import (
	"backend/models"
	"backend/services"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Helper functions to create pointers
func ptrString(s string) *string {
	return &s
}

func ptrBool(b bool) *bool {
	return &b
}

// GetTodosHandler のユニットテスト
func TestGetTodosHandler(t *testing.T) {
	// モックサービスのセットアップ
	mockService := new(services.MockTodoService)

	// テスト用のコントローラーを作成
	controller := NewRealTodoController(mockService)

	// 固定された日時
	fixedTime := models.NewCustomTime(time.Date(2024, time.September, 14, 0, 6, 27, 272515000, time.UTC))

	// テストデータの準備
	todos := []models.Todo{
		{
			ID:          "1",
			UserID:      "user1",
			Description: "Test Todo 1",
			Completed:   false,
			CreatedAt:   fixedTime,
			UpdatedAt:   fixedTime,
		},
		{
			ID:          "2",
			UserID:      "user2",
			Description: "Test Todo 2",
			Completed:   true,
			CreatedAt:   fixedTime,
			UpdatedAt:   fixedTime,
		},
	}

	// モックの期待値を設定
	mockService.On("GetAllTodos").Return(todos, nil)

	// Echoのインスタンスを作成
	e := echo.New()

	// リクエストとレスポンスの準備
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラーを実行
	if assert.NoError(t, controller.GetTodosHandler(c)) {
		// ステータスコードの検証
		assert.Equal(t, http.StatusOK, rec.Code)

		// レスポンスボディの検証
		var responseTodos []models.Todo
		if err := json.Unmarshal(rec.Body.Bytes(), &responseTodos); err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}

		assert.Equal(t, todos, responseTodos)
	}

	// モックの期待が満たされていることを確認
	mockService.AssertExpectations(t)
}

// CreateTodoHandler のユニットテスト
func TestCreateTodoHandler(t *testing.T) {
	// モックサービスのセットアップ
	mockService := new(services.MockTodoService)

	// テスト用のコントローラーを作成
	controller := NewRealTodoController(mockService)

	// 環境変数からUserIDを取得
	userID := os.Getenv("TEST_USER_ID")
	if userID == "" {
		t.Fatalf("USER_ID environment variable is not set")
	}

	// テストデータの準備
	reqBody := models.TodoCreateRequest{
		UserID:      userID,
		Description: "New Test Todo",
		Completed:   false,
	}

	// 固定された日時
	fixedTime := models.NewCustomTime(time.Date(2024, time.September, 14, 9, 6, 27, 272897000, time.UTC))

	todo := &models.Todo{
		ID:          "3",
		UserID:      reqBody.UserID,
		Description: reqBody.Description,
		Completed:   reqBody.Completed,
		CreatedAt:   fixedTime,
		UpdatedAt:   fixedTime,
	}

	// モックの期待値を設定
	mockService.On("CreateTodo", mock.AnythingOfType("*models.Todo")).Return(todo, nil)

	// Echoのインスタンスを作成
	e := echo.New()

	// リクエスト用のJSONデータを作成
	todoJSON, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// リクエストとレスポンスの準備
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラーを実行
	if assert.NoError(t, controller.CreateTodoHandler(c)) {
		// ステータスコードの検証
		assert.Equal(t, http.StatusCreated, rec.Code)

		// レスポンスボディの検証
		var responseTodo models.Todo
		if err := json.Unmarshal(rec.Body.Bytes(), &responseTodo); err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}

		assert.Equal(t, todo, &responseTodo)
	}

	// モックの期待が満たされていることを確認
	mockService.AssertExpectations(t)
}

// UpdateTodoHandler のユニットテスト
func TestUpdateTodoHandler(t *testing.T) {
	// モックサービスのセットアップ
	mockService := new(services.MockTodoService)

	// テスト用のコントローラーを作成
	controller := NewRealTodoController(mockService)

	// 環境変数からUserIDを取得
	userID := os.Getenv("TEST_USER_ID")
	if userID == "" {
		t.Fatalf("USER_ID environment variable is not set")
	}

	// テストデータの準備
	todoID := "1"
	updateReq := models.TodoUpdateRequest{
		UserID:      userID,
		Description: ptrString("Updated Test Todo"),
		Completed:   ptrBool(true),
	}

	// 固定された日時を設定
	fixedTime := models.NewCustomTime(time.Date(2024, time.September, 14, 9, 9, 22, 952253000, time.UTC))

	updatedTodo := &models.Todo{
		ID:          todoID,
		UserID:      "user1",
		Description: "Updated Test Todo",
		Completed:   true,
		CreatedAt:   fixedTime, // 固定日時を設定
		UpdatedAt:   fixedTime,
	}

	// モックの期待値を設定
	mockService.On("UpdateTodo", todoID, mock.AnythingOfType("*models.Todo")).Return(updatedTodo, nil)

	// Echoのインスタンスを作成
	e := echo.New()

	// 更新データをJSONに変換
	updateJSON, err := json.Marshal(updateReq)
	if err != nil {
		t.Fatalf("Failed to marshal update request: %v", err)
	}

	// リクエストとレスポンスの準備
	req := httptest.NewRequest(http.MethodPut, "/todos/"+todoID, bytes.NewReader(updateJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues(todoID)

	// ハンドラーを実行
	if assert.NoError(t, controller.UpdateTodoHandler(c)) {
		// ステータスコードの検証
		assert.Equal(t, http.StatusOK, rec.Code)

		// レスポンスボディの検証
		var responseTodo models.Todo
		if err := json.Unmarshal(rec.Body.Bytes(), &responseTodo); err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}

		assert.Equal(t, updatedTodo, &responseTodo)
	}

	// モックの期待が満たされていることを確認
	mockService.AssertExpectations(t)
}

// UpdateTodoHandler のユニットテスト（Todoが見つからない場合）
func TestUpdateTodoHandler_NotFound(t *testing.T) {
	// モックサービスのセットアップ
	mockService := new(services.MockTodoService)

	// テスト用のコントローラーを作成
	controller := NewRealTodoController(mockService)

	// 環境変数からUserIDを取得
	userID := os.Getenv("TEST_USER_ID")
	if userID == "" {
		t.Fatalf("USER_ID environment variable is not set")
	}

	// テストデータの準備
	todoID := "999"
	updateReq := models.TodoUpdateRequest{
		UserID:      userID,
		Description: ptrString("Non-existent Todo"),
		Completed:   ptrBool(true),
	}

	// モックの期待値を設定
	mockService.On("UpdateTodo", todoID, mock.AnythingOfType("*models.Todo")).Return((*models.Todo)(nil), services.ErrTodoNotFound)

	// Echoのインスタンスを作成
	e := echo.New()

	// 更新データをJSONに変換
	updateJSON, err := json.Marshal(updateReq)
	if err != nil {
		t.Fatalf("Failed to marshal update request: %v", err)
	}

	// リクエストとレスポンスの準備
	req := httptest.NewRequest(http.MethodPut, "/todos/"+todoID, bytes.NewReader(updateJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues(todoID)

	// ハンドラーを実行
	if assert.NoError(t, controller.UpdateTodoHandler(c)) {
		// ステータスコードの検証
		assert.Equal(t, http.StatusNotFound, rec.Code)

		// レスポンスボディの検証
		var response map[string]string
		if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}

		assert.Equal(t, "Todo not found", response["error"])
	}

	// モックの期待が満たされていることを確認
	mockService.AssertExpectations(t)
}

// UpdateTodoHandler のユニットテスト（サービスエラーの場合）
func TestUpdateTodoHandler_ServiceError(t *testing.T) {
	// モックサービスのセットアップ
	mockService := new(services.MockTodoService)

	// テスト用のコントローラーを作成
	controller := NewRealTodoController(mockService)

	// 環境変数からUserIDを取得
	userID := os.Getenv("TEST_USER_ID")
	if userID == "" {
		t.Fatalf("USER_ID environment variable is not set")
	}

	// テストデータの準備
	todoID := "1"
	updateReq := models.TodoUpdateRequest{
		UserID:      userID,
		Description: ptrString("Service Error Todo"),
		Completed:   ptrBool(true),
	}

	serviceError := errors.New("database connection failed")

	// モックの期待値を設定
	mockService.On("UpdateTodo", todoID, mock.AnythingOfType("*models.Todo")).Return((*models.Todo)(nil), serviceError)

	// Echoのインスタンスを作成
	e := echo.New()

	// 更新データをJSONに変換
	updateJSON, err := json.Marshal(updateReq)
	if err != nil {
		t.Fatalf("Failed to marshal update request: %v", err)
	}

	// リクエストとレスポンスの準備
	req := httptest.NewRequest(http.MethodPut, "/todos/"+todoID, bytes.NewReader(updateJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues(todoID)

	// ハンドラーを実行
	if assert.NoError(t, controller.UpdateTodoHandler(c)) {
		// ステータスコードの検証
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		// レスポンスボディの検証
		var response map[string]string
		if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}

		assert.Equal(t, "Failed to update todo", response["error"])
	}

	// モックの期待が満たされていることを確認
	mockService.AssertExpectations(t)
}

// DeleteTodoHandler のユニットテスト
func TestDeleteTodoHandler(t *testing.T) {
	// モックサービスのセットアップ
	mockService := new(services.MockTodoService)

	// テスト用のコントローラーを作成
	controller := NewRealTodoController(mockService)

	// テストデータの準備
	todoID := "1"

	// モックの期待値を設定
	mockService.On("DeleteTodo", todoID).Return(nil)

	// Echoのインスタンスを作成
	e := echo.New()

	// リクエストとレスポンスの準備
	req := httptest.NewRequest(http.MethodDelete, "/todos/"+todoID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues(todoID)

	// ハンドラーを実行
	if assert.NoError(t, controller.DeleteTodoHandler(c)) {
		// ステータスコードの検証
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}

	// モックの期待が満たされていることを確認
	mockService.AssertExpectations(t)
}

// DeleteTodoHandler のユニットテスト（Todoが見つからない場合）
func TestDeleteTodoHandler_NotFound(t *testing.T) {
	// モックサービスのセットアップ
	mockService := new(services.MockTodoService)

	// テスト用のコントローラーを作成
	controller := NewRealTodoController(mockService)

	// テストデータの準備
	todoID := "999"

	// モックの期待値を設定
	mockService.On("DeleteTodo", todoID).Return(services.ErrTodoNotFound)

	// Echoのインスタンスを作成
	e := echo.New()

	// リクエストとレスポンスの準備
	req := httptest.NewRequest(http.MethodDelete, "/todos/"+todoID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues(todoID)

	// ハンドラーを実行
	if assert.NoError(t, controller.DeleteTodoHandler(c)) {
		// ステータスコードの検証
		assert.Equal(t, http.StatusNotFound, rec.Code)

		// レスポンスボディの検証
		var response map[string]string
		if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}

		assert.Equal(t, "Todo not found", response["error"])
	}

	// モックの期待が満たされていることを確認
	mockService.AssertExpectations(t)
}

// DeleteTodoHandler のユニットテスト（サービスエラーの場合）
func TestDeleteTodoHandler_ServiceError(t *testing.T) {
	// モックサービスのセットアップ
	mockService := new(services.MockTodoService)

	// テスト用のコントローラーを作成
	controller := NewRealTodoController(mockService)

	// テストデータの準備
	todoID := "1"

	// サービスエラーの定義
	serviceError := errors.New("database connection failed")

	// モックの期待値を設定
	mockService.On("DeleteTodo", todoID).Return(serviceError)

	// Echoのインスタンスを作成
	e := echo.New()

	// リクエストとレスポンスの準備
	req := httptest.NewRequest(http.MethodDelete, "/todos/"+todoID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues(todoID)

	// ハンドラーを実行
	if assert.NoError(t, controller.DeleteTodoHandler(c)) {
		// ステータスコードの検証
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		// レスポンスボディの検証
		var response map[string]string
		if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}

		assert.Equal(t, "Failed to delete todo", response["error"])
	}

	// モックの期待が満たされていることを確認
	mockService.AssertExpectations(t)
}

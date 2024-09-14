package controllers

import (
	"backend/models"
	"backend/services"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// 実際のTodoコントローラー
type RealTodoController struct {
	todoService services.TodoService
}

// RealControllerのコンストラクタ
func NewRealTodoController(todoService services.TodoService) *RealTodoController {
	return &RealTodoController{
		todoService: todoService,
	}
}

// 全てのTodoを取得するハンドラー
func (c *RealTodoController) GetTodosHandler(ctx echo.Context) error {
	todos, err := c.todoService.GetAllTodos()
	if err != nil {
		ctx.Logger().Errorf("Failed to get todos: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch todos"})
	}

	return ctx.JSON(http.StatusOK, todos)
}

// 新しいTodoを作成するハンドラー
func (c *RealTodoController) CreateTodoHandler(ctx echo.Context) error {
	var req models.TodoCreateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	/// バリデーション
	if req.UserID == "" || req.Description == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Validation failed"})
	}

	// Todoオブジェクトの作成
	todo := &models.Todo{
		UserID:      req.UserID,
		Description: req.Description,
		Completed:   req.Completed,
	}
	todo.GenerateID()
	todo.CreatedAt = models.NewCustomTime(time.Now())
	todo.UpdatedAt = models.NewCustomTime(time.Now())

	newTodo, err := c.todoService.CreateTodo(todo)
	if err != nil {
		ctx.Logger().Errorf("Failed to create todo: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create todo"})
	}

	return ctx.JSON(http.StatusCreated, newTodo)
}

// Todoを更新するハンドラー
func (c *RealTodoController) UpdateTodoHandler(ctx echo.Context) error {
	id := ctx.Param("id")

	var req models.TodoUpdateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// バリデーション
	if req.Description == nil || req.UserID == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Validation failed"})
	}

	// 更新データの準備
	todo := &models.Todo{
		ID:        id,
		UpdatedAt: models.NewCustomTime(time.Now()),
	}

	if req.Description != nil {
		todo.Description = *req.Description
	}
	if req.Completed != nil {
		todo.Completed = *req.Completed
	}
	if req.UserID != "" {
		todo.UserID = req.UserID
	}

	updatedTodo, err := c.todoService.UpdateTodo(id, todo)
	if err != nil {
		ctx.Logger().Errorf("Failed to update todo with ID %s: %v", id, err)
		if errors.Is(err, services.ErrTodoNotFound) {
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update todo"})
	}

	// updatedTodoがnilの場合の処理
	if updatedTodo == nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update todo"})
	}

	return ctx.JSON(http.StatusOK, updatedTodo)
}

// Todoを削除するハンドラー
func (c *RealTodoController) DeleteTodoHandler(ctx echo.Context) error {
	id := ctx.Param("id")

	err := c.todoService.DeleteTodo(id)
	if err != nil {
		ctx.Logger().Errorf("Failed to delete todo with ID %s: %v", id, err)
		if errors.Is(err, services.ErrTodoNotFound) {
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete todo"})
	}

	return ctx.NoContent(http.StatusNoContent)
}

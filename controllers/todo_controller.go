package controllers

import (
	"github.com/labstack/echo/v4"
)

// Todoコントローラーのインターフェース
type TodoController interface {
	CreateTodoHandler(echo.Context) error
	UpdateTodoHandler(echo.Context) error
	GetTodosHandler(echo.Context) error
	DeleteTodoHandler(echo.Context) error
}

package controllers

import (
	"github.com/labstack/echo/v4"
)

// Todoコントローラーのインターフェース
type TodoController interface {
	GetTodosHandler(echo.Context) error
	CreateTodoHandler(echo.Context) error
	UpdateTodoHandler(echo.Context) error
	DeleteTodoHandler(echo.Context) error
}

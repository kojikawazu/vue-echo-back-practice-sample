package controllers

import (
	"backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 全てのTodoを取得する
func GetTodos(c echo.Context) error {
	todos, err := services.GetAllTodos()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error fetching todos")
	}

	return c.JSON(http.StatusOK, todos)
}

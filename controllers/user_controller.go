package controllers

import (
	"backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetUsers は全てのユーザーを取得する
func GetUsers(c echo.Context) error {
	users, err := services.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error fetching users")
	}
	return c.JSON(http.StatusOK, users)
}

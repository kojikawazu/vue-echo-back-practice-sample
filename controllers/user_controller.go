package controllers

import (
	"github.com/labstack/echo/v4"
)

// Userコントローラーのインターフェース
type UserController interface {
	GetAllUsersHandler(c echo.Context) error
}

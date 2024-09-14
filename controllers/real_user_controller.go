package controllers

import (
	"backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 実際のUserコントローラー
type RealUserController struct {
	userService services.UserService
}

// RealUserControllerのコンストラクタ
func NewRealUserController(userService services.UserService) *RealUserController {
	return &RealUserController{
		userService: userService,
	}
}

// 全てのユーザーを取得する
func (c *RealUserController) GetAllUsersHandler(ctx echo.Context) error {
	users, err := c.userService.GetAllUsers()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error fetching users")
	}

	return ctx.JSON(http.StatusOK, users)
}

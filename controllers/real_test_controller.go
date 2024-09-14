package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// 実際のTestコントローラー
type RealTestController struct {
}

// RealTestControllerのコンストラクタ
func NewRealTestController() *RealTestController {
	return &RealTestController{}
}

/**
 * テスト用のハンドラー
 */
func (c *RealTestController) TestHandler(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}

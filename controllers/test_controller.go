package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/**
 * テスト用のハンドラー
 */
func TestHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

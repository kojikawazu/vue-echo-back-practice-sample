package controllers

import "github.com/labstack/echo/v4"

// TestControllerのインターフェース
type TestController interface {
	TestHandler(c echo.Context) error
}

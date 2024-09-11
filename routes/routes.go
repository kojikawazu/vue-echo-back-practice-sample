package routes

import (
	"backend/controllers"

	"github.com/labstack/echo/v4"
)

/**
 * ルーティングの設定
 */
func InitRoutes(e *echo.Echo) {
	e.GET("/test", controllers.TestHandler)
}

package routes

import (
	"backend/controllers"

	"github.com/labstack/echo/v4"
)

/**
 * ルーティングの設定
 */
func InitRoutes(e *echo.Echo) {
	e.GET("/", controllers.TestHandler)

	e.GET("/users", controllers.GetUsers)

	e.GET("/todos", controllers.GetTodos)
}

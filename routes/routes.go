package routes

import (
	"backend/controllers"

	"github.com/labstack/echo/v4"
)

// Echoのルーターにハンドラーを登録
func SetupRoutes(e *echo.Echo, todoController controllers.Controller) {

	//e.GET("/", controllers.TestHandler)
	//e.GET("/users", controllers.GetUsers)

	// ルートグループの作成
	todos := e.Group("/todos")

	// ハンドラーの登録
	todos.GET("", todoController.GetTodosHandler)
	todos.POST("", todoController.CreateTodoHandler)
	todos.PUT("/:id", todoController.UpdateTodoHandler)
	todos.DELETE("/:id", todoController.DeleteTodoHandler)
}

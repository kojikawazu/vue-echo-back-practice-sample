package routes

import (
	"backend/controllers"

	"github.com/labstack/echo/v4"
)

// Echoのルーターにハンドラーを登録
func SetupRoutes(e *echo.Echo, todoController controllers.TodoController, testController controllers.TestController) {

	//e.GET("/users", controllers.GetUsers)

	// ルートグループの作成
	todos := e.Group("/todos")
	test := e.Group("/test")

	test.GET("", testController.TestHandler)

	// ハンドラーの登録
	todos.GET("", todoController.GetTodosHandler)
	todos.POST("", todoController.CreateTodoHandler)
	todos.PUT("/:id", todoController.UpdateTodoHandler)
	todos.DELETE("/:id", todoController.DeleteTodoHandler)
}

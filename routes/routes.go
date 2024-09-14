package routes

import (
	"backend/controllers"

	"github.com/labstack/echo/v4"
)

// Echoのルーターにハンドラーを登録
func SetupRoutes(e *echo.Echo, todoController controllers.TodoController, testController controllers.TestController, userController controllers.UserController) {

	// ルートグループの作成
	todos := e.Group("/todos")
	test := e.Group("/test")
	user := e.Group("/users")

	// ハンドラーの登録
	test.GET("", testController.TestHandler)

	user.GET("", userController.GetAllUsersHandler)

	todos.GET("", todoController.GetTodosHandler)
	todos.POST("", todoController.CreateTodoHandler)
	todos.PUT("/:id", todoController.UpdateTodoHandler)
	todos.DELETE("/:id", todoController.DeleteTodoHandler)
}

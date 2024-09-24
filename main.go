package main

import (
	"backend/controllers"
	"backend/lib"
	"backend/routes"
	"backend/services"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// 環境変数の読み込み
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	// Supabaseクライアントの初期化
	client := lib.InitSupabaseClient()

	// サービスの初期化
	todoService := services.NewRealTodoService(client)
	userService := services.NewRealUserService(client)

	// コントローラーの初期化
	todoController := controllers.NewRealTodoController(todoService)
	userController := controllers.NewRealUserController(userService)
	testController := controllers.NewRealTestController()

	// Echoのインスタンスを作成
	e := echo.New()

	// CORSを有効化
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("FRONTEND_URL")},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// ルーティングの設定
	routes.SetupRoutes(e, todoController, testController, userController)

	// サーバーの起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}

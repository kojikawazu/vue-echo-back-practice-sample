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

	// コントローラーの初期化
	todoController := controllers.NewRealController(todoService)

	// Echoのインスタンスを作成
	e := echo.New()

	// ルーティングの設定
	routes.SetupRoutes(e, todoController)

	// サーバーの起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}

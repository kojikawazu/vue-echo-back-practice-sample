package main

import (
	"backend/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	// .envファイルをロード
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティングの設定
	routes.InitRoutes(e)

	// サーバーの開始
	e.Logger.Fatal(e.Start(":8080"))
}

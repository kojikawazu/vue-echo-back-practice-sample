package main

import (
	"backend/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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

package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    // ルートハンドラー
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, Echo!")
    })

    // サーバーの開始
    e.Logger.Fatal(e.Start(":8080"))
}

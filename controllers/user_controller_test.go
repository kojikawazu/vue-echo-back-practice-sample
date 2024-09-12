package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// テスト環境用の.env.testファイルを読み込む
	err := godotenv.Load("../.env.test")
	if err != nil {
		panic("Error loading ../.env.test file")
	}

	// テストを実行
	code := m.Run()

	// 終了コードを渡して終了
	os.Exit(code)
}

func TestUserHandler(t *testing.T) {
	// Echoのインスタンスを作成
	e := echo.New()

	// `/users`ルートのテスト
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// GetUsersを実行し、レスポンスの確認
	if assert.NoError(t, GetUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "user")
	}
}

package main

import (
	"backend/routes"
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

func setupServer() *httptest.Server {
	// Echoのインスタンスを作成
	e := echo.New()

	// ルーティングを初期化
	routes.InitRoutes(e)

	// 実際にサーバーを起動し、テストサーバーとして扱う
	server := httptest.NewServer(e.Server.Handler)
	return server
}

func TestE2EHelloWorld(t *testing.T) {
	// テストサーバーをセットアップ
	server := setupServer()
	defer server.Close()

	// サーバーへのリクエストを作成
	resp, err := http.Get(server.URL + "/")
	assert.NoError(t, err)
	defer resp.Body.Close()

	// ステータスコードを確認
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// レスポンスボディを確認
	body := make([]byte, 1024)
	n, _ := resp.Body.Read(body)
	assert.Contains(t, string(body[:n]), "Hello, World!")
}

func TestE2EGetUsers(t *testing.T) {
	// テストサーバーをセットアップ
	server := setupServer()
	defer server.Close()

	// `/users`エンドポイントへのリクエストを作成
	resp, err := http.Get(server.URL + "/users")
	assert.NoError(t, err)
	defer resp.Body.Close()

	// ステータスコードを確認
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// レスポンスボディを確認
	body := make([]byte, 1024)
	n, _ := resp.Body.Read(body)
	assert.Contains(t, string(body[:n]), "username")
}

func TestE2EGetTodos(t *testing.T) {
	// テストサーバーをセットアップ
	server := setupServer()
	defer server.Close()

	// `/todos`エンドポイントへのリクエストを作成
	resp, err := http.Get(server.URL + "/todos")
	assert.NoError(t, err)
	defer resp.Body.Close()

	// ステータスコードを確認
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// レスポンスボディを確認
	body := make([]byte, 1024)
	n, _ := resp.Body.Read(body)
	assert.Contains(t, string(body[:n]), "description")
}

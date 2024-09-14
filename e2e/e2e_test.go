package main

import (
	"backend/controllers"
	"backend/lib"
	"backend/routes"
	"backend/services"
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

func setupServer() *echo.Echo {
	// Supabaseクライアントの初期化
	client := lib.InitSupabaseClient()

	// サービスの初期化
	todoService := services.NewRealTodoService(client)

	// コントローラーの初期化
	todoController := controllers.NewRealController(todoService)

	// Echoのインスタンスを作成
	e := echo.New()

	// ルーティングの設定（コントローラーを渡す）
	routes.SetupRoutes(e, todoController)

	return e
}

// func TestE2EHelloWorld(t *testing.T) {
// 	// テストサーバーをセットアップ
// 	server := setupServer()
// 	defer server.Close()

// 	// サーバーへのリクエストを作成
// 	resp, err := http.Get(server.URL + "/")
// 	assert.NoError(t, err)
// 	defer resp.Body.Close()

// 	// ステータスコードを確認
// 	assert.Equal(t, http.StatusOK, resp.StatusCode)

// 	// レスポンスボディを確認
// 	body := make([]byte, 1024)
// 	n, _ := resp.Body.Read(body)
// 	assert.Contains(t, string(body[:n]), "Hello, World!")
// }

// func TestE2EGetUsers(t *testing.T) {
// 	// テストサーバーをセットアップ
// 	server := setupServer()
// 	defer server.Close()

// 	// `/users`エンドポイントへのリクエストを作成
// 	resp, err := http.Get(server.URL + "/users")
// 	assert.NoError(t, err)
// 	defer resp.Body.Close()

// 	// ステータスコードを確認
// 	assert.Equal(t, http.StatusOK, resp.StatusCode)

// 	// レスポンスボディを確認
// 	body := make([]byte, 1024)
// 	n, _ := resp.Body.Read(body)
// 	assert.Contains(t, string(body[:n]), "username")
// }

func TestE2EGetTodos(t *testing.T) {
	// サーバーをセットアップ
	e := setupServer()

	// リクエストを作成
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラーを実行
	todoController := controllers.NewRealController(services.NewRealTodoService(lib.InitSupabaseClient()))
	if assert.NoError(t, todoController.GetTodosHandler(c)) {
		// ステータスコードを確認
		assert.Equal(t, http.StatusOK, rec.Code)

		// レスポンスボディを確認
		assert.Contains(t, rec.Body.String(), "description")
	}
}

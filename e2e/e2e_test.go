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
	userService := services.NewRealUserService(client)

	// コントローラーの初期化
	todoController := controllers.NewRealTodoController(todoService)
	userController := controllers.NewRealUserController(userService)
	testController := controllers.NewRealTestController()

	// Echoのインスタンスを作成
	e := echo.New()

	// ルーティングの設定（コントローラーを渡す）
	routes.SetupRoutes(e, todoController, testController, userController)

	return e
}

func TestE2EHelloWorld(t *testing.T) {
	// サーバーをセットアップ
	e := setupServer()

	// リクエストを作成
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラーを実行
	testController := controllers.NewRealTestController()
	if assert.NoError(t, testController.TestHandler(c)) {
		// ステータスコードを確認
		assert.Equal(t, http.StatusOK, rec.Code)

		// レスポンスボディを確認
		assert.Contains(t, rec.Body.String(), "Hello, World!")
	}
}

func TestE2EGetAllUsers(t *testing.T) {
	// サーバーをセットアップ
	e := setupServer()

	// リクエストを作成
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラーを実行
	userController := controllers.NewRealUserController(services.NewRealUserService(lib.InitSupabaseClient()))
	if assert.NoError(t, userController.GetAllUsersHandler(c)) {
		// ステータスコードを確認
		assert.Equal(t, http.StatusOK, rec.Code)

		// レスポンスボディを確認
		assert.Contains(t, rec.Body.String(), "username")
	}
}

func TestE2EGetTodos(t *testing.T) {
	// サーバーをセットアップ
	e := setupServer()

	// リクエストを作成
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラーを実行
	todoController := controllers.NewRealTodoController(services.NewRealTodoService(lib.InitSupabaseClient()))
	if assert.NoError(t, todoController.GetTodosHandler(c)) {
		// ステータスコードを確認
		assert.Equal(t, http.StatusOK, rec.Code)

		// レスポンスボディを確認
		assert.Contains(t, rec.Body.String(), "description")
	}
}

package routes

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
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

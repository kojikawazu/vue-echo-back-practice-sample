package lib

import (
	"log"
	"os"
	"sync"

	"github.com/go-resty/resty/v2"
)

// Supabase APIクライアントのグローバル変数
var supabaseClient *resty.Client
var once sync.Once // 初期化を一度だけ行うためにsync.Onceを使用

// Supabaseクライアントを初期化
func InitSupabaseClient() *resty.Client {
	once.Do(func() {
		// 環境変数を取得
		supabaseURL := os.Getenv("SUPABASE_URL")
		supabaseKey := os.Getenv("SUPABASE_KEY")

		// SupabaseのURLとキーが正しく設定されているか確認
		if supabaseURL == "" || supabaseKey == "" {
			log.Fatal("SUPABASE_URL or SUPABASE_KEY is not set")
		}

		// Supabase APIクライアントの作成
		supabaseClient = resty.New().
			SetBaseURL(supabaseURL).
			SetHeader("apikey", supabaseKey).
			SetHeader("Authorization", "Bearer "+supabaseKey).
			SetHeader("Content-Type", "application/json")
	})

	return supabaseClient
}

package services

import (
	"backend/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/go-resty/resty/v2"
)

// Supabase APIクライアントのグローバル変数
var supabaseClient *resty.Client
var once sync.Once // 初期化を一度だけ行うためにsync.Onceを使用

// Supabaseクライアントを初期化
func initSupabaseClient() {
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
}

// Supabaseから全てのユーザーを取得する
func GetAllUsers() ([]models.User, error) {
	// Supabaseクライアントを初期化（遅延初期化）
	initSupabaseClient()

	// APIリクエストを送信
	resp, err := supabaseClient.R().
		Get("/rest/v1/users?select=*")

	if err != nil {
		fmt.Println("Error making request to Supabase:", err)
		return nil, err
	}

	// レスポンスステータスの確認
	if resp.StatusCode() != 200 {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode())
		fmt.Println("Response body:", string(resp.Body()))
		return nil, fmt.Errorf("failed to fetch users, status code: %d", resp.StatusCode())
	}

	// レスポンスをパースして、ユーザーデータを取得
	var users []models.User
	if err := json.Unmarshal(resp.Body(), &users); err != nil {
		fmt.Println("Error parsing response:", err)
		return nil, err
	}
	return users, nil
}

package services

import (
	"backend/lib"
	"backend/models"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// 実際のTodoサービス
type RealUserService struct {
	Client *resty.Client
}

// RealUserServiceのコンストラクタ
func NewRealUserService(client *resty.Client) *RealUserService {
	return &RealUserService{Client: client}
}

// Supabaseから全てのユーザーを取得する
func (s *RealUserService) GetAllUsers() ([]models.User, error) {
	// Supabaseクライアントを初期化
	client := lib.InitSupabaseClient()

	// APIリクエストを送信
	resp, err := client.R().
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

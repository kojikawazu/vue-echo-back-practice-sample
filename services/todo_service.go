package services

import (
	"backend/lib"
	"backend/models"
	"encoding/json"
	"fmt"
)

// Supabaseから全てのTodoを取得する
func GetAllTodos() ([]models.Todo, error) {
	// Supabaseクライアントを初期化
	client := lib.InitSupabaseClient()

	// APIリクエストを送信
	resp, err := client.R().
		Get("/rest/v1/todos?select=*")

	if err != nil {
		fmt.Println("Error making request to Supabase:", err)
		return nil, err
	}

	// レスポンスステータスの確認
	if resp.StatusCode() != 200 {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode())
		fmt.Println("Response body:", string(resp.Body()))
		return nil, fmt.Errorf("failed to fetch todos, status code: %d", resp.StatusCode())
	}

	// レスポンスをパースして、Todoデータを取得
	var todos []models.Todo
	if err := json.Unmarshal(resp.Body(), &todos); err != nil {
		fmt.Println("Error parsing response:", err)
		return nil, err
	}
	return todos, nil
}

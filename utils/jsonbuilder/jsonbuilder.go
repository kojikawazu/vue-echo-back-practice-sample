package jsonbuilder

import (
	"encoding/json"
	"fmt"
	"os"
)

// Todo は Todo 構造体
type Todo struct {
	UserID      string `json:"user_id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// NewTodoJSON は Todo の JSON データを作成する関数
func NewTodoJSON(description string, completed bool) (string, error) {
	// 環境変数から user_id を取得
	userID := os.Getenv("TEST_USER_ID")
	if userID == "" {
		return "", fmt.Errorf("user_id is not set")
	}

	todo := Todo{
		UserID:      userID,
		Description: description,
		Completed:   completed,
	}

	// JSON に変換
	todoJSON, err := json.Marshal(todo)
	if err != nil {
		return "", fmt.Errorf("error marshalling todo: %v", err)
	}

	return string(todoJSON), nil
}

package models

import (
	"github.com/google/uuid"
)

type Todo struct {
	ID          string     `json:"id" db:"id"`                   // UUID型
	UserID      string     `json:"user_id" db:"user_id"`         // ユーザーID(UUID型)
	Description string     `json:"description" db:"description"` // 説明
	Completed   bool       `json:"completed" db:"completed"`     // 完了フラグ
	CreatedAt   CustomTime `json:"created_at" db:"created_at"`   // カスタムフォーマットで時間をパース
	UpdatedAt   CustomTime `json:"updated_at" db:"updated_at"`   // カスタムフォーマットで時間をパース
}

// TodoCreateRequest はTodo作成時のリクエスト構造体
type TodoCreateRequest struct {
	UserID      string `json:"user_id" validate:"required,uuid"`
	Description string `json:"description" validate:"required"`
	Completed   bool   `json:"completed"`
}

// TodoUpdateRequest はTodo更新時のリクエスト構造体
type TodoUpdateRequest struct {
	UserID      string  `json:"user_id" validate:"required,uuid"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}

// TodoのIDを生成する関数
func (todo *Todo) GenerateID() {
	todo.ID = uuid.NewString()
}

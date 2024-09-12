package models

// Supabaseのusersテーブルに対応するモデル
type User struct {
	ID        string     `json:"id" db:"id"`                 // UUID型
	Username  string     `json:"username" db:"username"`     // ユーザー名
	Password  string     `json:"-" db:"password"`            // パスワードはレスポンスに含めない
	Email     string     `json:"email" db:"email"`           // メールアドレス
	CreatedAt CustomTime `json:"created_at" db:"created_at"` // カスタムフォーマットで時間をパース
	UpdatedAt CustomTime `json:"updated_at" db:"updated_at"` // カスタムフォーマットで時間をパース
}

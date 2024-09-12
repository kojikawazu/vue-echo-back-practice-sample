package models

import "time"

// カスタムの時間フォーマットを扱うための型
type CustomTime struct {
	time.Time
}

// Supabaseから返されるカスタム時間フォーマットを解析する
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	// 時間フォーマットを明示的に指定する（タイムゾーンなしのフォーマット）
	const layout = "2006-01-02T15:04:05.999999"
	parsedTime, err := time.Parse(`"`+layout+`"`, string(b))

	if err != nil {
		return err
	}

	ct.Time = parsedTime
	return nil
}

// Supabaseのusersテーブルに対応するモデル
type User struct {
	ID        string     `json:"id" db:"id"`                 // UUID型
	Username  string     `json:"username" db:"username"`     // ユーザー名
	Password  string     `json:"-" db:"password"`            // パスワードはレスポンスに含めない
	Email     string     `json:"email" db:"email"`           // メールアドレス
	CreatedAt CustomTime `json:"created_at" db:"created_at"` // カスタムフォーマットで時間をパース
	UpdatedAt CustomTime `json:"updated_at" db:"updated_at"` // カスタムフォーマットで時間をパース
}

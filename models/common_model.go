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

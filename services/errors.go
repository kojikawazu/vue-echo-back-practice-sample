package services

import "errors"

// Todoが見つからなかった場合のエラー
var ErrTodoNotFound = errors.New("todo not found")

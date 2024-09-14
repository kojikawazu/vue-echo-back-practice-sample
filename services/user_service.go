package services

import "backend/models"

// USerサービスのインターフェース
type UserService interface {
	GetAllUsers() ([]models.User, error)
}

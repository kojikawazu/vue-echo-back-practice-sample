package services

import (
	"backend/models"

	"github.com/stretchr/testify/mock"
)

// UserService インターフェースのモック
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetAllUsers() ([]models.User, error) {
	args := m.Called()
	// 型アサーションの安全性を確保
	if users, ok := args.Get(0).([]models.User); ok {
		return users, args.Error(1)
	}
	return nil, args.Error(1)
}

package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

// UserController インターフェースのモック
type MockUserController struct {
	mock.Mock
}

func (m *MockUserController) GetAllUsersHandler(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

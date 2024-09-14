package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

// TodoController インターフェースのモック
type MockTodoController struct {
	mock.Mock
}

func (m *MockTodoController) GetTodosHandler(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockTodoController) CreateTodoHandler(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockTodoController) UpdateTodoHandler(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockTodoController) DeleteTodoHandler(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

package services

import (
	"backend/models"

	"github.com/stretchr/testify/mock"
)

// TodoService インターフェースのモック
type MockTodoService struct {
	mock.Mock
}

func (m *MockTodoService) GetAllTodos() ([]models.Todo, error) {
	args := m.Called()
	return args.Get(0).([]models.Todo), args.Error(1)
}

func (m *MockTodoService) CreateTodo(todo *models.Todo) (*models.Todo, error) {
	args := m.Called(todo)
	return args.Get(0).(*models.Todo), args.Error(1)
}

func (m *MockTodoService) UpdateTodo(id string, todo *models.Todo) (*models.Todo, error) {
	args := m.Called(id, todo)
	return args.Get(0).(*models.Todo), args.Error(1)
}

func (m *MockTodoService) DeleteTodo(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

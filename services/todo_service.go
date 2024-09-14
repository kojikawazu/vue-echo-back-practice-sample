package services

import "backend/models"

// Todoサービスのインターフェース
type TodoService interface {
	GetAllTodos() ([]models.Todo, error)
	CreateTodo(todo *models.Todo) (*models.Todo, error)
	UpdateTodo(id string, todo *models.Todo) (*models.Todo, error)
	DeleteTodo(id string) error
}

package services

import (
	"backend/models"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// 実際のTodoサービス
type RealTodoService struct {
	Client *resty.Client
}

// RealTodoServiceのコンストラクタ
func NewRealTodoService(client *resty.Client) *RealTodoService {
	return &RealTodoService{Client: client}
}

// Supabaseから全てのTodoを取得する
func (s *RealTodoService) GetAllTodos() ([]models.Todo, error) {
	resp, err := s.Client.R().
		Get("/rest/v1/todos?select=*")
	if err != nil {
		fmt.Println("Error making request to Supabase:", err)
		return nil, err
	}

	if resp.StatusCode() != 200 {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode())
		fmt.Println("Response body:", string(resp.Body()))
		return nil, fmt.Errorf("failed to fetch todos, status code: %d", resp.StatusCode())
	}

	var todos []models.Todo
	if err := json.Unmarshal(resp.Body(), &todos); err != nil {
		fmt.Println("Error parsing response:", err)
		return nil, err
	}

	return todos, nil
}

// Supabaseに新しいTodoを作成
func (s *RealTodoService) CreateTodo(todo *models.Todo) (*models.Todo, error) {
	todoJSON, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling todo:", err)
		return nil, err
	}

	resp, err := s.Client.R().
		SetHeader("Prefer", "return=representation").
		SetBody(todoJSON).
		Post("/rest/v1/todos")
	if err != nil {
		fmt.Println("Error creating todo in Supabase:", err)
		return nil, err
	}

	// ステータスコードのログを出力
	//fmt.Printf("Supabase Response Status: %d\n", resp.StatusCode())
	//fmt.Println("Supabase Response Body:", string(resp.Body()))

	if resp.StatusCode() != 201 {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode())
		fmt.Println("Response body:", string(resp.Body()))
		return nil, fmt.Errorf("failed to create todo, status code: %d", resp.StatusCode())
	}

	var createdTodos []models.Todo
	if err := json.Unmarshal(resp.Body(), &createdTodos); err != nil {
		fmt.Println("Error parsing response:", err)
		return nil, err
	}

	if len(createdTodos) > 0 {
		return &createdTodos[0], nil
	}

	return nil, fmt.Errorf("no todo created")
}

// Supabase上のTodoを更新
func (s *RealTodoService) UpdateTodo(id string, todo *models.Todo) (*models.Todo, error) {
	todoJSON, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling todo:", err)
		return nil, err
	}

	resp, err := s.Client.R().
		SetHeader("Prefer", "return=representation").
		SetBody(todoJSON).
		Put(fmt.Sprintf("/rest/v1/todos?id=eq.%s", id))
	if err != nil {
		fmt.Println("Error updating todo in Supabase:", err)
		return nil, err
	}

	// ステータスコードのログを出力
	//fmt.Printf("Supabase Response Status: %d\n", resp.StatusCode())
	//fmt.Println("Supabase Response Body:", string(resp.Body()))

	if resp.StatusCode() == 404 {
		return nil, ErrTodoNotFound
	}

	if resp.StatusCode() != 200 {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode())
		fmt.Println("Response body:", string(resp.Body()))
		return nil, fmt.Errorf("failed to update todo, status code: %d", resp.StatusCode())
	}

	// レスポンスが配列として返されるため、配列でパースする
	var updatedTodos []models.Todo
	if err := json.Unmarshal(resp.Body(), &updatedTodos); err != nil {
		fmt.Println("Error parsing response:", err)
		return nil, err
	}

	// 更新されたTodoが配列の最初の要素として返されるので、それを返す
	if len(updatedTodos) > 0 {
		return &updatedTodos[0], nil
	}

	return nil, fmt.Errorf("no todo updated")
}

// Supabase上のTodoを削除
func (s *RealTodoService) DeleteTodo(id string) error {
	resp, err := s.Client.R().
		SetHeader("Prefer", "return=representation").
		Delete(fmt.Sprintf("/rest/v1/todos?id=eq.%s", id))
	if err != nil {
		fmt.Println("Error deleting todo in Supabase:", err)
		return err
	}

	if resp.StatusCode() == 404 {
		return ErrTodoNotFound
	}

	if resp.StatusCode() != 200 {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode())
		fmt.Println("Response body:", string(resp.Body()))
		return fmt.Errorf("failed to delete todo, status code: %d", resp.StatusCode())
	}

	return nil
}

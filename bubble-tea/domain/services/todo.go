package services

import (
	"bubbletea/domain/models"
)

// TodoService Serving todo data
type TodoService struct {
}

// ListTodos lists all todos
func (srv *TodoService) ListTodos() ([]models.Todo, error) {
	return []models.Todo{
		{
			Id:      0,
			Title:   "Buy Milk",
			Details: "Need to buy milk for coffee",
		},
		{
			Id:      1,
			Title:   "Mop House",
			Details: "Time to mop the house",
		},
	}, nil
}

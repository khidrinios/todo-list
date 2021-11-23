package service

import "khidr/todo/persistence"

type Service struct {
	postgres *persistence.PostgresConfig
}

func New(postgres *persistence.PostgresConfig) *Service {
	return &Service{
		postgres: postgres,
	}
}

func (s *Service) CreateTodo(title string, description *string) (*int, error) {
	todoId, err := s.postgres.CreateTodo(title, description)
	if err != nil {
		return nil, err
	}
	return todoId, nil
}

func (s *Service) GetTodoById(id int) (*persistence.Todo, error) {
	todo, err := s.postgres.GetTodoById(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *Service) UpdateTodo(id int, title, description *string, isDone *bool) (*persistence.Todo, error) {
	todo, err := s.postgres.UpdateTodo(id, title, description, isDone)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

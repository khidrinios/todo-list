package todo

import (
	"khidr/todo/interfaces/api/todo"
	"khidr/todo/persistence"
)

type Service struct {
	postgres *persistence.PostgresConfig
}

func NewService(postgres *persistence.PostgresConfig) *Service {
	return &Service{
		postgres: postgres,
	}
}

func (s *Service) CreateTodo(req todo.CreateTodoRequestBody) (*todo.TodoIdResult, error) {
	todoId, err := s.postgres.CreateTodo(req.Title, req.Description)
	if err != nil {
		return nil, err
	}
	return &todo.TodoIdResult{Id: *todoId}, nil
}

func (s *Service) GetTodoById(req todo.TodoByIdRequestUri) (*todo.TodoResult, error) {
	todoRes, err := s.postgres.GetTodoById(req.ID)
	if err != nil {
		return nil, err
	}
	return &todo.TodoResult{
		CreatedAt:   todoRes.CreatedAt,
		Description: todoRes.Description,
		Id:          int(todoRes.ID),
		IsDone:      todoRes.IsDone,
		Title:       todoRes.Title,
		UpdatedAt:   todoRes.UpdatedAt,
	}, nil
}

func (s *Service) QueryTodos(req todo.QueryTodosRequestBody) ([]todo.TodoResult, error) {
	offset := getOffsetFromPageAndLimit(req.Page, req.Limit)
	todos, err := s.postgres.QueryTodos(req.Title, req.Description, req.IsDone, offset, req.Limit)
	if err != nil {
		return nil, err
	}
	todosRes := make([]todo.TodoResult, 0)
	for i := range todos {
		res := todos[i]
		todoRes := todo.TodoResult{
			CreatedAt:   res.CreatedAt,
			Description: res.Description,
			Id:          int(res.ID),
			IsDone:      res.IsDone,
			Title:       res.Title,
			UpdatedAt:   res.UpdatedAt,
		}
		todosRes = append(todosRes, todoRes)
	}
	return todosRes, nil
}

func getOffsetFromPageAndLimit(page, limit int) int {
	offset := (page - 1) * limit
	return offset
}

func (s *Service) DeleteTodoById(req todo.TodoByIdRequestUri) (*todo.TodoIdResult, error) {
	todoId, err := s.postgres.DeleteTodoById(req.ID)
	if err != nil {
		return nil, err
	}
	return &todo.TodoIdResult{
		Id: *todoId,
	}, nil
}

func (s *Service) UpdateTodo(reqParam todo.TodoByIdRequestUri, reqbody todo.UpdateTodoRequestBody) (*todo.UpdateTodoResult, error) {
	todoRes, err := s.postgres.UpdateTodo(reqParam.ID, reqbody.Title, reqbody.Description, reqbody.IsDone)
	if err != nil {
		return nil, err
	}
	return &todo.UpdateTodoResult{
		CreatedAt:   todoRes.CreatedAt,
		Description: todoRes.Description,
		Id:          int(todoRes.ID),
		IsDone:      todoRes.IsDone,
		Title:       todoRes.Title,
		UpdatedAt:   todoRes.UpdatedAt,
	}, nil
}

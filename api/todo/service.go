package todo

import (
	iterface "khidr/todo/interfaces/api/todo"
	"khidr/todo/models"
)

type Service struct {
	repo iterface.Repository
}

func NewService(repo iterface.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(req iterface.CreateRequest) (*iterface.TodoIdResult, error) {
	todo := models.Todo{
		Title:       req.Title,
		Description: req.Description,
		IsDone:      false,
	}
	todoId, err := s.repo.Create(todo)
	if err != nil {
		return nil, err
	}
	return &iterface.TodoIdResult{Id: *todoId}, nil
}

func (s *Service) GetById(req iterface.TodoByIdRequestUri) (*models.Todo, error) {
	return s.repo.GetTodoById(req.ID)
}

func (s *Service) List(req iterface.QueryTodosRequestBody) ([]models.Todo, error) {
	offset := getOffsetFromPageAndLimit(req.Page, req.Limit)
	repoRequest := models.QueryTodosRequest{
		Title:       req.Title,
		Description: req.Description,
		IsDone:      req.IsDone,
		Offset:      offset,
		Limit:       req.Limit,
	}
	todos, err := s.repo.QueryTodos(repoRequest)
	return todos, err
}

func getOffsetFromPageAndLimit(page, limit int) int {
	offset := (page - 1) * limit
	return offset
}

func (s *Service) DeleteById(req iterface.TodoByIdRequestUri) (*iterface.TodoIdResult, error) {
	todoId, err := s.repo.DeleteTodoById(req.ID)
	if err != nil {
		return nil, err
	}
	return &iterface.TodoIdResult{
		Id: *todoId,
	}, nil
}

func (s *Service) Update(reqParam iterface.TodoByIdRequestUri, reqbody iterface.UpdateTodoRequestBody) (*iterface.UpdateTodoResult, error) {
	todoRes, err := s.repo.UpdateTodo(reqParam.ID, reqbody.Title, reqbody.Description, reqbody.IsDone)
	if err != nil {
		return nil, err
	}
	return &iterface.UpdateTodoResult{
		CreatedAt:   todoRes.CreatedAt,
		Description: todoRes.Description,
		Id:          int(todoRes.ID),
		IsDone:      todoRes.IsDone,
		Title:       todoRes.Title,
		UpdatedAt:   todoRes.UpdatedAt,
	}, nil
}

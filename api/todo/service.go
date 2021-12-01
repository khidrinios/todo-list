package todo

import (
	"fmt"
	todoI "khidr/todo/interfaces/api/todo"
	"khidr/todo/models"
	"time"
)

type Service struct {
	repo todoI.Repository
}

func NewService(repo todoI.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

type ValidationError struct {
	err error
}

func (v ValidationError) Error() string {
	return v.err.Error()
}

func (s *Service) Create(req todoI.CreateRequest) (*int, error) {
	from, to := GetBeginningAndEndTime(time.Now())
	queryReq := todoI.QueryTodosRequest{
		Title: &req.Title,
		From:  &from,
		To:    &to,
	}
	existingTodos, err := s.repo.QueryTodos(queryReq)
	if err != nil {
		return nil, err
	}
	if len(existingTodos) > 0 {
		err := fmt.Errorf("%s already exists", req.Title)
		validationError := ValidationError{err: err}
		return nil, validationError
	}

	todo := models.Todo{
		Title:       req.Title,
		Description: req.Description,
		IsDone:      false,
	}
	todoId, err := s.repo.Create(todo)
	return todoId, err
}

func GetBeginningAndEndTime(t time.Time) (time.Time, time.Time) {
	year, month, day := t.Date()
	from := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	to := from.Add(24 * time.Hour)
	return from, to
}

func (s *Service) GetById(req todoI.TodoByIdRequestUri) (*models.Todo, error) {
	return s.repo.GetTodoById(req.ID)
}

func (s *Service) List(req todoI.ListRequestBody) ([]models.Todo, error) {
	offset := getOffsetFromPageAndLimit(req.Page, req.Limit)
	repoRequest := todoI.QueryTodosRequest{
		Title:       req.Title,
		Description: req.Description,
		IsDone:      req.IsDone,
		Offset:      offset,
		Limit:       req.Limit,
	}
	return s.repo.QueryTodos(repoRequest)
}

func getOffsetFromPageAndLimit(page, limit int) int {
	offset := (page - 1) * limit
	return offset
}

func (s *Service) DeleteById(req todoI.TodoByIdRequestUri) (*int, error) {
	return s.repo.DeleteTodoById(req.ID)
}

func (s *Service) Update(reqParam todoI.TodoByIdRequestUri, reqbody todoI.UpdateTodoRequestBody) (*models.Todo, error) {
	return s.repo.UpdateTodo(reqParam.ID, reqbody.Title, reqbody.Description, reqbody.IsDone)
}

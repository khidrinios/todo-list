package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"gorm.io/gorm"
)

func (v *V1) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var body CreateTodoRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		handleError(w, r, err, err.Error(), http.StatusBadRequest)
		return
	}

	if body.Title == "" {
		handleError(w, r, nil, "Title must not be empty", http.StatusBadRequest)
		return
	}

	todoId, err := v.svc.CreateTodo(body.Title, body.Description)
	if err != nil {
		handleError(w, r, err, err.Error(), http.StatusInternalServerError)
		return
	}

	res := CreateTodoResult{
		Id: *todoId,
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, res)
}

func (v *V1) GetTodoById(w http.ResponseWriter, r *http.Request, id int) {
	if id < 1 {
		handleError(w, r, nil, "Todo Id must be greater than zero", http.StatusBadRequest)
		return
	}

	todo, err := v.svc.GetTodoById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handleError(w, r, err, err.Error(), http.StatusNotFound)
			return
		}
		handleError(w, r, err, err.Error(), http.StatusInternalServerError)
		return
	}

	res := TodoResult{
		CreatedAt:   todo.CreatedAt,
		Description: todo.Description,
		Id:          int(todo.ID),
		IsDone:      todo.IsDone,
		Title:       todo.Title,
		UpdatedAt:   todo.UpdatedAt,
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

func (v *V1) UpdateTodo(w http.ResponseWriter, r *http.Request, id int) {
	if id < 1 {
		handleError(w, r, nil, "Todo Id must be greater than zero", http.StatusBadRequest)
		return
	}
	var body UpdateTodoRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		handleError(w, r, err, err.Error(), http.StatusBadRequest)
		return
	}
	if body.Title != nil && len(*body.Title) == 0 {
		handleError(w, r, nil, "Title cannot be empty", http.StatusBadRequest)
		return
	}
	todo, err := v.svc.UpdateTodo(id, body.Title, body.Description, body.IsDone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handleError(w, r, err, err.Error(), http.StatusNotFound)
			return
		}
		handleError(w, r, err, err.Error(), http.StatusInternalServerError)
		return
	}

	res := UpdateTodoResult{
		CreatedAt:   todo.CreatedAt,
		Description: todo.Description,
		Id:          int(todo.ID),
		IsDone:      todo.IsDone,
		Title:       todo.Title,
		UpdatedAt:   todo.UpdatedAt,
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)

}

func (v *V1) QueryTodos(w http.ResponseWriter, r *http.Request) {
	var body QueryTodosRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		handleError(w, r, err, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := v.svc.QueryTodos(body.Title, body.Description, body.IsDone, body.Page, body.Limit)
	if err != nil {
		handleError(w, r, err, err.Error(), http.StatusInternalServerError)
		return
	}

	todos := make([]TodoResult, 0)
	for i := range res {
		todoRes := res[i]
		todo := TodoResult{
			CreatedAt:   todoRes.CreatedAt,
			Description: todoRes.Description,
			Id:          int(todoRes.ID),
			IsDone:      todoRes.IsDone,
			Title:       todoRes.Title,
			UpdatedAt:   todoRes.UpdatedAt,
		}
		todos = append(todos, todo)
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, todos)
}

func (v *V1) DeleteTodoById(w http.ResponseWriter, r *http.Request, id int) {
	if id < 1 {
		handleError(w, r, nil, "Todo Id must be greater than zero", http.StatusBadRequest)
		return
	}
	deletedTodoId, err := v.svc.DeleteTodoById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			handleError(w, r, err, err.Error(), http.StatusNotFound)
			return
		}
		handleError(w, r, err, err.Error(), http.StatusInternalServerError)
		return
	}

	res := map[string]interface{}{
		"id": *deletedTodoId,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

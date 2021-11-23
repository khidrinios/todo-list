package v1

import (
	"encoding/json"
	"khidr/todo/service"
	"net/http"

	"github.com/go-chi/render"
	"github.com/labstack/gommon/log"
)

type V1 struct {
	svc *service.Service
}

func NewV1(svc *service.Service) *V1 {
	return &V1{
		svc: svc,
	}
}

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

func handleError(w http.ResponseWriter, r *http.Request, err error, message string, status int) {
	if message == "" {
		message = "Internal Error"
	}
	errors := Error{Errors: []string{message}}
	log.Error(err)
	render.Status(r, status)
	render.JSON(w, r, errors)
}

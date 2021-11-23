package v1

import (
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

func handleError(w http.ResponseWriter, r *http.Request, err error, message string, status int) {
	if message == "" {
		message = "Internal Error"
	}
	error := Error{Error: message}
	log.Error(err)
	render.Status(r, status)
	render.JSON(w, r, error)
}

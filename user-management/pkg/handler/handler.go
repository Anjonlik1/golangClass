package handler

import (
	"user-management/pkg/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	services *service.Service
}

var response = map[string]any{}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", h.users)
	r.HandleFunc("/users/{id}", h.singleUser)
	r.HandleFunc("/users", h.createUser)

	return r
}

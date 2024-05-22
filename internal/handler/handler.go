package handler

import (
	"net/http"
	"pastebin/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /pastebin/", h.create)
	mux.HandleFunc("GET /pastebin/{id}", h.get)

	return mux
}
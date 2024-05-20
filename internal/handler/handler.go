package handler

import "net/http"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /log-in", h.logIn)
	mux.HandleFunc("POST /sign-up", h.signUp)

	mux.HandleFunc("POST /", h.create)
	mux.HandleFunc("GET /{id}", h.get)
	mux.HandleFunc("PATCH /{id}", h.update)
	mux.HandleFunc("DELETE /{id}", h.delete)

	return mux
}
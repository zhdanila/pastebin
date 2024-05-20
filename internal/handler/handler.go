package handler

import "net/http"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /pastebin/", h.create)
	mux.HandleFunc("GET /pastebin/{id}", h.get)
	mux.HandleFunc("DELETE /pastebin/{id}", h.delete)

	return mux
}
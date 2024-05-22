package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pastebin/internal/models"
)

func(h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var userPaste models.UserPaste

	err := json.NewDecoder(r.Body).Decode(&userPaste)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Create(userPaste)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("id - %s", id)))
}

func(h *Handler) get(w http.ResponseWriter, r *http.Request) {
	pasteId := r.PathValue("id")

	paste, err := h.services.Paste.Get(pasteId, "")
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	marshalledPaste, err := json.Marshal(&paste)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledPaste)
}


func(h *Handler) delete(w http.ResponseWriter, r *http.Request) {

}
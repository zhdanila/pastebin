package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/logger"
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
	if _, err := w.Write([]byte(fmt.Sprintf("id - %s", id))); err != nil {
		// Якщо виникла помилка під час запису відповіді, логуємо її
		logger.Errorf("Failed to write response: %v", err)
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}

func(h *Handler) get(w http.ResponseWriter, r *http.Request) {
	var passwordInput models.PasswordInput

	err := json.NewDecoder(r.Body).Decode(&passwordInput)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	pasteId := r.PathValue("id")

	paste, err := h.services.Paste.Get(pasteId, passwordInput.Password)
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
	if _, err := w.Write(marshalledPaste); err != nil {
		// Якщо виникла помилка під час запису відповіді, логуємо її
		logger.Errorf("Failed to write response: %v", err)
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}



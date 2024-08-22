package handler

import (
	"github.com/google/logger"
	"net/http"
)

func NewErrorResponse(w http.ResponseWriter, statusCode int, errorString string) {
	logger.Error(errorString)
	w.WriteHeader(statusCode)
	if _, err := w.Write([]byte(errorString)); err != nil {
		// Якщо виникла помилка під час запису відповіді, логуємо її
		logger.Errorf("Failed to write response: %v", err)
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}
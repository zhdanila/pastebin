package handler

import (
	"github.com/google/logger"
	"net/http"
)

func NewErrorResponse(w http.ResponseWriter, statusCode int, errorString string) {
	logger.Error(errorString)
	w.WriteHeader(statusCode)
	w.Write([]byte(errorString))
}
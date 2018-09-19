package utils

import (
	"encoding/json"
	"net/http"
)

// HTTP constants
const (
	ApplicaitonJSON = "application/json"
	ContentType     = "Content-Type"
)

// Success - Response formatter for Successful requests
func Success(w http.ResponseWriter, entity interface{}) {
	w.Header().Set(ContentType, ApplicaitonJSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entity)
}

// Error - Response formatter for Error conditions
func Error(w http.ResponseWriter, entity interface{}, httpStatus int) {
	w.Header().Set(ContentType, ApplicaitonJSON)
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(entity)
}

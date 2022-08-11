package handler

import (
	"encoding/json"
	"net/http"
)

type Failure struct {
	Message string `json:"message"`
}

func (f Failure) returnJSON(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(f)
}

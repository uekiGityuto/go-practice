package handler

import (
	"encoding/json"
	"net/http"
)

type Failure struct {
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}

func (f Failure) returnJSON(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if f.Detail == nil {
		f.Detail = ""
	}
	json.NewEncoder(w).Encode(f)
}

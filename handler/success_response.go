package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Success struct {
	Message string `json:"message"`
}

func (s Success) returnJSON(w http.ResponseWriter) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	if err := encoder.Encode(s); err != nil {
		log.Println("Error:", err)
		f := Failure{
			Message: "システムエラーが発生しました",
		}
		f.returnJSON(w, http.StatusInternalServerError)
	}
	if _, err := fmt.Fprint(w, buf.String()); err != nil {
		log.Println("Error:", err)
		f := Failure{
			Message: "システムエラーが発生しました",
		}
		f.returnJSON(w, http.StatusInternalServerError)
		return
	}
}

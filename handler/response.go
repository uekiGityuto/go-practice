package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/uekiGityuto/go-practice/usecase"
	"golang.org/x/xerrors"
	"log"
	"net/http"
)

func returnResponse(w http.ResponseWriter, body interface{}) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	if err := encoder.Encode(body); err != nil {
		err = xerrors.Errorf("レスポンスボディに書き込む情報のJSONシリアライズが失敗しました。: %w", err)
		log.Printf("Error: %+v\n", err)
		return
	}
	if _, err := fmt.Fprint(w, buf.String()); err != nil {
		err = xerrors.Errorf("レスポンスボディへの書き込みが失敗しました。: %w", err)
		log.Printf("Error: %+v\n", err)
		return
	}
}

type ErrorResponse struct {
	Message string            `json:"message"`
	Detail  map[string]string `json:"detail,omitempty"`
}

func returnError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err.(type) {
	case ValidationError:
		w.WriteHeader(http.StatusBadRequest)
		validationErr := err.(ValidationError)
		body := ErrorResponse{
			Message: validationErr.Error(),
			Detail:  validationErr.detail,
		}
		returnResponse(w, body)
		return
	case NotFound:
		w.WriteHeader(http.StatusNotFound)
		notFoundErr := err.(NotFound)
		body := ErrorResponse{
			Message: notFoundErr.Error(),
		}
		returnResponse(w, body)
		return
	}

	switch {
	case errors.Is(err, usecase.NotFoundErr):
		w.WriteHeader(http.StatusBadRequest)
		body := ErrorResponse{Message: usecase.NotFoundErr.Error()}
		returnResponse(w, body)
		return
	default:
		log.Printf("Error: %+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		body := ErrorResponse{Message: "システムエラーです。"}
		returnResponse(w, body)
		return
	}
}

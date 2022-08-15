package ui

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

func ReturnResponse(w http.ResponseWriter, body interface{}) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	if err := encoder.Encode(body); err != nil {
		err = xerrors.Errorf("レスポンスボディに書き込む情報のJSONシリアライズが失敗しました。: %w", err)
		log.Printf("%+v\n", err)
		return
	}
	if _, err := fmt.Fprint(w, buf.String()); err != nil {
		err = xerrors.Errorf("レスポンスボディへの書き込みが失敗しました。: %w", err)
		log.Printf("%+v\n", err)
		return
	}
}

type ErrorResponse struct {
	Message string            `json:"message"`
	Detail  map[string]string `json:"detail,omitempty"`
}

func ReturnError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// handlerで発生したユーザ定義エラーのハンドリング
	var validationErr ValidationError
	var notFound NotFound
	switch {
	case errors.As(err, &validationErr):
		w.WriteHeader(http.StatusBadRequest)
		body := ErrorResponse{
			Message: validationErr.Error(),
			Detail:  validationErr.Detail,
		}
		ReturnResponse(w, body)
		return
	case errors.As(err, &notFound):
		w.WriteHeader(http.StatusNotFound)
		body := ErrorResponse{
			Message: notFound.Error(),
		}
		ReturnResponse(w, body)
		return
	}

	// usecaseで発生したユーザ定義エラーのハンドリング
	switch {
	case errors.Is(err, usecase.ErrNotFound):
		w.WriteHeader(http.StatusBadRequest)
		body := ErrorResponse{Message: usecase.ErrNotFound.Error()}
		ReturnResponse(w, body)
		return
	}

	// システムエラーのハンドリング
	log.Printf("%+v\n", err)
	w.WriteHeader(http.StatusInternalServerError)
	body := ErrorResponse{Message: "システムエラーです。"}
	ReturnResponse(w, body)
}

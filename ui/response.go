package ui

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/uekiGityuto/go-practice/lib/log"
	"github.com/uekiGityuto/go-practice/usecase"
	"go.uber.org/zap"
	"net/http"
)

func ReturnResponse(w http.ResponseWriter, body interface{}) {
	var buf bytes.Buffer
	logger := log.Logger
	encoder := json.NewEncoder(&buf)
	if err := encoder.Encode(body); err != nil {
		logger.Error("レスポンスボディに書き込む情報のJSONシリアライズが失敗しました。", zap.Error(err))
		return
	}
	if _, err := fmt.Fprint(w, buf.String()); err != nil {
		logger.Error("レスポンスボディへの書き込みが失敗しました。", zap.Error(err))
		return
	}
}

type ErrorResponse struct {
	Message string            `json:"message"`
	Detail  map[string]string `json:"detail,omitempty"`
}

func ReturnError(w http.ResponseWriter, err error, msg string) {
	logger := log.Logger
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
	logger.Error(msg, zap.Error(err))
	w.WriteHeader(http.StatusInternalServerError)
	body := ErrorResponse{Message: "システムエラーです。"}
	ReturnResponse(w, body)
}

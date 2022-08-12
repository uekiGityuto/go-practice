package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/xerrors"
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
		err = xerrors.Errorf("レスポンスボディに書き込む内容のJSONシリアライズが失敗しました。: %w", err)
		log.Printf("Error: %+v\n", err)
		f := Failure{
			Message: "システムエラーが発生しました",
		}
		f.returnJSON(w, http.StatusInternalServerError)
	}
	if _, err := fmt.Fprint(w, buf.String()); err != nil {
		err = xerrors.Errorf("レスポンスボディへの書き込みが失敗しました。: %w", err)
		log.Printf("Error: %+v\n", err)
		f := Failure{
			Message: "システムエラーが発生しました",
		}
		f.returnJSON(w, http.StatusInternalServerError)
		return
	}
}

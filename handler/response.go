package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/xerrors"
	"log"
	"net/http"
)

func createResponse(w http.ResponseWriter, body interface{}) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	if err := encoder.Encode(body); err != nil {
		err = xerrors.Errorf("レスポンスボディに書き込む情報のJSONシリアライズが失敗しました。: %w", err)
		log.Printf("Error: %+v\n", err)
		f := Failure{
			Message: "システムエラーが発生しました",
		}
		f.returnJSON(w, http.StatusInternalServerError)
		return
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

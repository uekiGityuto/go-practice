package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type User struct {
	ID         string `json:"id"`
	FamilyName string `json:"family_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	Sex        string `json:"sex"`
}

var UserHandler = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch strings.ToUpper(r.Method) {
	case http.MethodGet:
		get(w, r)
	case http.MethodPost:
		post(w, r)
	default:
		f := Failure{
			Message: "サポートされていないHTTPメソッドです。",
		}
		f.returnJSON(w, http.StatusNotFound)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		f := Failure{
			Message: "パラメータが不正です。",
		}
		f.returnJSON(w, http.StatusBadRequest)
		return
	}
	user := User{
		ID:         id,
		FamilyName: "ueki",
		LastName:   "yuto",
		Age:        28,
		Sex:        "男",
	}

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	if err := encoder.Encode(user); err != nil {
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

func post(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&user); err != nil {
		log.Println("Error:", err)
		f := Failure{
			Message: "システムエラーが発生しました",
		}
		f.returnJSON(w, http.StatusInternalServerError)
	}

	fmt.Printf("response body: %+v\n", user)

	s := Success{Message: "正常に登録しました。"}
	s.returnJSON(w)
}

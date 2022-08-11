package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/uekiGityuto/go-practice/usecase"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type User struct {
	UseCase usecase.User
}

func NewUser(uc usecase.User) *User {
	return &User{
		UseCase: uc,
	}
}

type Form struct {
	FamilyName string `json:"family_name" validate:"required"`
	GivenName  string `json:"given_name" validate:"required"`
	Age        int    `json:"age" validate:"required,gte=0"`
	Sex        string `json:"sex" validate:"required"`
}

func (form *Form) validate() (ok bool, result map[string]string) {
	result = make(map[string]string)
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	if err := validate.Struct(form); err != nil {
		ok = false
		errors := err.(validator.ValidationErrors)
		for _, err := range errors {
			switch err.StructField() {
			case "FamilyName":
				result[err.Field()] = err.Tag()
			case "GivenName":
				result[err.Field()] = err.Tag()
			case "Age":
				result[err.Field()] = err.Tag()
			case "Sex":
				result[err.Field()] = err.Tag()
			}
		}
	} else {
		ok = true
	}
	return ok, result
}

type UserResponse struct {
	ID         string `json:"id"`
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
	Age        int    `json:"age"`
	Sex        string `json:"sex"`
}

func (h User) HandleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch strings.ToUpper(r.Method) {
	case http.MethodGet:
		h.get(w, r)
	case http.MethodPost:
		h.post(w, r)
	default:
		f := Failure{
			Message: "サポートされていないHTTPメソッドです。",
		}
		f.returnJSON(w, http.StatusNotFound)
	}
}

func (h User) get(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		f := Failure{
			Message: "パラメータが不正です。",
		}
		f.returnJSON(w, http.StatusBadRequest)
		return
	}

	entity, err := h.UseCase.Find(id)
	if err != nil {
		log.Println("Error:", err)
		f := Failure{
			Message: "システムエラーが発生しました",
		}
		f.returnJSON(w, http.StatusInternalServerError)
		return
	}

	user := UserResponse{
		ID:         entity.ID.String(),
		FamilyName: entity.FamilyName,
		GivenName:  entity.GivenName,
		Age:        entity.Age,
		Sex:        entity.Sex,
	}

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	if err := encoder.Encode(user); err != nil {
		log.Println("Error:", err)
		f := Failure{
			Message: "システムエラーが発生しました",
		}
		f.returnJSON(w, http.StatusInternalServerError)
		return
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

func (h User) post(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var form Form
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&form); err != nil {
		log.Println("Error:", err)
		f := Failure{
			Message: "システムエラーが発生しました",
		}
		f.returnJSON(w, http.StatusInternalServerError)
		return
	}
	if ok, result := form.validate(); !ok {
		f := Failure{
			Message: "パラメータが不正です。",
			Detail:  result,
		}
		f.returnJSON(w, http.StatusBadRequest)
		return
	}

	//fmt.Printf("response body: %+v\n", form)
	if err := h.UseCase.Save(form.FamilyName, form.GivenName, form.Age, form.Sex); err != nil {
		log.Println("Error:", err)
		f := Failure{
			Message: "システムエラーが発生しました",
		}
		f.returnJSON(w, http.StatusInternalServerError)
		return
	}

	s := Success{Message: "正常に登録しました。"}
	s.returnJSON(w)
}

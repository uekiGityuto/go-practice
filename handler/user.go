package handler

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/uekiGityuto/go-practice/usecase"
	"golang.org/x/xerrors"
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
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
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

type GetResponse struct {
	ID         string `json:"id"`
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
	Age        int    `json:"age"`
	Sex        string `json:"sex"`
}

type PostResponse struct {
	ID string `json:"id"`
}

func (h *User) HandleUser(w http.ResponseWriter, r *http.Request) {
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

func (h *User) get(w http.ResponseWriter, r *http.Request) {
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
		if errors.Is(err, usecase.ErrNotFound) {
			f := Failure{
				Message: usecase.ErrNotFound.Error(),
			}
			f.returnJSON(w, http.StatusBadRequest)
			return
		}
		err = xerrors.Errorf("ユーザ情報取得が失敗しました。: %w", err)
		log.Printf("Error: %+v\n", err)
		f := Failure{
			Message: "システムエラーが発生しました",
		}
		f.returnJSON(w, http.StatusInternalServerError)
		return
	}

	createResponse(w, GetResponse{
		ID:         entity.ID.String(),
		FamilyName: entity.FamilyName,
		GivenName:  entity.GivenName,
		Age:        entity.Age,
		Sex:        entity.Sex,
	})
}

func (h *User) post(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var form Form
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&form); err != nil {
		err = xerrors.Errorf("リクエストボディのJSONデシリアライズが失敗しました: %w", err)
		log.Printf("Error: %+v\n", err)
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
	id, err := h.UseCase.Save(form.FamilyName, form.GivenName, form.Age, form.Sex)
	if err != nil {
		err = xerrors.Errorf("ユーザ登録が失敗しました。: %w", err)
		log.Printf("Error: %+v\n", err)
		f := Failure{
			Message: "システムエラーが発生しました",
		}
		f.returnJSON(w, http.StatusInternalServerError)
		return
	}

	createResponse(w, PostResponse{
		ID: id.String(),
	})
}

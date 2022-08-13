package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/uekiGityuto/go-practice/ui"
	customValidator "github.com/uekiGityuto/go-practice/ui/validator"
	"github.com/uekiGityuto/go-practice/usecase"
	"golang.org/x/xerrors"
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

type GetForm struct {
	ID string `json:"id" validator:"required,uuid4"`
}

func (form *GetForm) validate() error {
	validate := customValidator.Get()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	detail := make(map[string]string)
	if err := validate.Struct(form); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			switch err.StructField() {
			case "ID":
				detail[err.Field()] = err.Tag()
			}
		}
		return ui.NewValidationError(detail)
	} else {
		return nil
	}
}

type PostForm struct {
	FamilyName string `json:"family_name" validator:"required"`
	GivenName  string `json:"given_name" validator:"required"`
	Age        int    `json:"age" validator:"required,gte=0"`
	Sex        string `json:"sex" validator:"required,sex"`
}

func (form *PostForm) validate() error {
	validate := customValidator.Get()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	detail := make(map[string]string)
	if err := validate.Struct(form); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			switch err.StructField() {
			case "FamilyName":
				detail[err.Field()] = err.Tag()
			case "GivenName":
				detail[err.Field()] = err.Tag()
			case "Age":
				detail[err.Field()] = err.Tag()
			case "Sex":
				detail[err.Field()] = err.Tag()
			}
		}
		return ui.NewValidationError(detail)
	} else {
		return nil
	}
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
		err := ui.NotFound{Message: "サポートされていないHTTPメソッドです。"}
		ui.ReturnError(w, err)
	}
}

func (h *User) get(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	form := GetForm{ID: id}
	if err := form.validate(); err != nil {
		ui.ReturnError(w, err)
		return
	}

	entity, err := h.UseCase.Find(id)
	if err != nil {
		err = xerrors.Errorf("ユーザ情報取得が失敗しました。: %w", err)
		ui.ReturnError(w, err)
		return
	}

	ui.ReturnResponse(w, GetResponse{
		ID:         entity.ID.String(),
		FamilyName: entity.FamilyName,
		GivenName:  entity.GivenName,
		Age:        entity.Age,
		Sex:        entity.Sex,
	})
}

func (h *User) post(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var form PostForm
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&form); err != nil {
		err = xerrors.Errorf("リクエストボディのJSONデシリアライズが失敗しました: %w", err)
		ui.ReturnError(w, err)
		return
	}
	if err := form.validate(); err != nil {
		ui.ReturnError(w, err)
		return
	}
	id, err := h.UseCase.Save(form.FamilyName, form.GivenName, form.Age, form.Sex)
	if err != nil {
		err = xerrors.Errorf("ユーザ登録が失敗しました。: %w", err)
		ui.ReturnError(w, err)
		return
	}

	ui.ReturnResponse(w, PostResponse{
		ID: id.String(),
	})
}

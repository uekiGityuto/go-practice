package validator

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/xerrors"
)

var validate = validator.New()

func validateSex(fl validator.FieldLevel) bool {
	typeList := []string{
		"男",
		"女",
	}
	for _, v := range typeList {
		if v == fl.Field().String() {
			return true
		}
	}
	return false
}

func RegisterCustomValidator() error {
	if err := validate.RegisterValidation("sex", validateSex); err != nil {
		return xerrors.Errorf("カスタムバリデーションの登録に失敗しました。: %w", err)
	}
	return nil
}

func Get() *validator.Validate {
	return validate
}

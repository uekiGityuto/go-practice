package ui

type ValidationError struct {
	Message string
	Detail  map[string]string
}

func (err ValidationError) Error() string {
	return err.Message
}

func NewValidationError(detail map[string]string) ValidationError {
	return ValidationError{
		Message: "パラメーターが不正です。",
		Detail:  detail,
	}
}

type NotFound struct {
	Message string
}

func (err NotFound) Error() string {
	return err.Message
}

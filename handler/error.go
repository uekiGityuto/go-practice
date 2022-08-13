package handler

type ValidationError struct {
	message string
	detail  map[string]string
}

func (err ValidationError) Error() string {
	return err.message
}

func NewValidationError(detail map[string]string) ValidationError {
	return ValidationError{
		message: "パラメーターが不正です。",
		detail:  detail,
	}
}

type NotFound struct {
	message string
}

func (err NotFound) Error() string {
	return err.message
}

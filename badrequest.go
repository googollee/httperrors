package httperrors

import (
	"bytes"
	"fmt"
	"net/http"
)

type FieldError[Code CodeType] struct {
	Name    string `json:"name"`
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func Field[Code CodeType](name string, code Code, message string) FieldError[Code] {
	return FieldError[Code]{
		Name:    name,
		Code:    code,
		Message: message,
	}
}

type BadRequestError[Code CodeType] struct {
	Message string                      `json:"message"`
	Details map[string]FieldError[Code] `json:"details"`
}

func BadRequest[Code CodeType](message string, fields ...FieldError[Code]) *BadRequestError[Code] {
	ret := &BadRequestError[Code]{
		Message: message,
		Details: make(map[string]FieldError[Code]),
	}

	return ret.Join(fields...)
}

func (e *BadRequestError[Code]) Join(fields ...FieldError[Code]) *BadRequestError[Code] {
	for _, field := range fields {
		e.Details[field.Name] = field
	}
	return e
}

func (e BadRequestError[Code]) HTTPCode() int {
	return http.StatusBadRequest
}

func (e BadRequestError[Code]) Error() string {
	var buf bytes.Buffer

	for _, field := range e.Details {
		fmt.Fprintf(&buf, "%s: (%s)%s, ", field.Name, field.Code, field.Message)
	}

	return fmt.Sprintf("(%d)%s: %s", e.HTTPCode(), e.Message, buf.String())
}

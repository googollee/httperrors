package httperrors

import "fmt"

type HTTPCode interface {
	HTTPCode() int
}

type CodeType interface {
	~string
}

type HTTPError[Code CodeType] struct {
	httpCode int

	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func Error[Code CodeType](httpCode int, code Code, message string) *HTTPError[Code] {
	return &HTTPError[Code]{
		httpCode: httpCode,
		Code:     code,
		Message:  message,
	}
}

func (e HTTPError[Code]) HTTPCode() int {
	return e.httpCode
}

func (e HTTPError[Code]) Error() string {
	return fmt.Sprintf("(%d:%s)%s", e.HTTPCode(), e.Code, e.Message)
}

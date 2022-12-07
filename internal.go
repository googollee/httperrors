package httperrors

import "net/http"

func InternalError[Code CodeType](err error) *HTTPError[Code] {
	return &HTTPError[Code]{
		httpCode: http.StatusInternalServerError,
		Message:  err.Error(),
	}
}

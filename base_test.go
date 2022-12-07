package httperrors

import (
	"net/http"
	"testing"
)

func TestBaseAsError(t *testing.T) {
	err := Error(http.StatusBadRequest, "bad_request", "reason")
	var _ error = err
	var _ HTTPCode = err
}


package httperrors

import (
	"io"
	"testing"
)

func TestInternalAsError(t *testing.T) {
	err := InternalError[string](io.EOF)
	var _ error = err
	var _ HTTPCode = err
}

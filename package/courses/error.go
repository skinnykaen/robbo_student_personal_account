package courses

import "errors"

var (
	ErrBadRequest          = errors.New("bad request")
	ErrInternalServerLevel = errors.New("internal server level error")
	ErrBadRequestBody      = errors.New("bad request body")
)

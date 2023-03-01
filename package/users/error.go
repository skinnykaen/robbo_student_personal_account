package users

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrBadRequest          = errors.New("bad request")
	ErrBadRequestBody      = errors.New("bad request body")
	ErrAlreadyUsedEmail    = errors.New("email is already used by another user")
)

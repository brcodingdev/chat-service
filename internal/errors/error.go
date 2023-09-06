package errors

import (
	"errors"
)

var (
	ErrInvalidLogin = errors.New("invalid login")
	ErrInRequest    = errors.New("bad request")
	ErrDupEmail     = errors.New("email already exists")
	ErrToken        = errors.New("error jwt token")
)

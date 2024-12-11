package errs

import "errors"

var (
	UserNotFound       = errors.New("user doesn't exist")
	ErrInvalidPassword = errors.New("invalid password provided")
)

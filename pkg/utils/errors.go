package errors

import "errors"

var (
	ErrInvalidIdFormat   = errors.New("invalid id format")
	ErrUserAlreadyExists = errors.New("user already exists")
)

package errors

import "errors"

var (
	ErrInvalidIdFormat     = errors.New("invalid id format")
	ErrEntityAlreadyExists = errors.New("entity already exists")
)

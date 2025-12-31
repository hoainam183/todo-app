package apperrors

import "errors"

var (
	ErrUserExists   = errors.New("user already exists")
	ErrMissingField = errors.New("missing field")
)

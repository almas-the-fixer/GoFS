package apperrors

import "errors"

var (
	ErrDuplicateEmail = errors.New("Email already Exists!")
	ErrDuplicateUsername = errors.New("Username already Exists!")
	ErrUniqueConstraintViolated = errors.New("SQL Unique Constraint Violated!")
)
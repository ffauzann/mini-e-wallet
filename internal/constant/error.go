package constant

import (
	"errors"
)

// Known errors
var (
	// Generic errors
	ErrBadRequest          = errors.New("Bad request")
	ErrUnauthorized        = errors.New("Unauthorized")
	ErrNotFound            = errors.New("Not found")
	ErrUnprocessableEntity = errors.New("Unprocessable entity")
	ErrInternal            = errors.New("Internal error")

	// Specific errors
	ErrAccountAlreadyExists   = errors.New("Account already exists")
	ErrAccountAlreadyEnabled  = errors.New("Account already enabled")
	ErrAccountAlreadyDisabled = errors.New("Account already disabled")
	ErrAccountDisabled        = errors.New("Account disabled")
	ErrDuplicateReferenceID   = errors.New("Duplicate reference ID")
	ErrInsufficientBalance    = errors.New("Insufficient balance")
)

var (
	MapErrorToStatusCode = map[error]int{
		// Generic errors
		ErrBadRequest:          400,
		ErrUnauthorized:        401,
		ErrNotFound:            404,
		ErrUnprocessableEntity: 422,
		ErrInternal:            500,

		// Specific errors
		ErrAccountAlreadyExists:   400,
		ErrAccountAlreadyEnabled:  400,
		ErrAccountAlreadyDisabled: 400,
		ErrAccountDisabled:        400,
		ErrDuplicateReferenceID:   400,
		ErrInsufficientBalance:    400,
	}
)

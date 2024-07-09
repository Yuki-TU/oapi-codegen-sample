package myerrors

import (
	"github.com/cockroachdb/errors"
)

var (
	ErrNotUser    = errors.New("does not exist user")
	ErrUnexpected = errors.New("unexpected error")
)

package model

import "errors"

var (
	ErrDuplicateEmail   = errors.New("duplicate email")
	ErrRecordNotFound   = errors.New("record not found")
	ErrPasswordMismatch = errors.New("password mismatch")
)

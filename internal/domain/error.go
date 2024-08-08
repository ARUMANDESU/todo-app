package domain

import "errors"

var (
	ErrTaskNotFound     = errors.New("task not found")
	ErrInvalidArguments = errors.New("invalid arguments")
	ErrInternal         = errors.New("internal error")
	ErrCancelled        = errors.New("cancelled")
)

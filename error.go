package collection

import "errors"

var (
	// ErrInvalidIndex indicates the given index is invalid for the operation.
	ErrInvalidIndex = errors.New("invalid index")

	// ErrHeadNotFound indicates the head node is not found.
	ErrHeadNotFound = errors.New("head not found")

	// ErrTailNotFound indicates the tail node is not found.
	ErrTailNotFound = errors.New("tail not found")
)

package errs

import "errors"

const (
	RESOURCE_NOT_FOUND = "resource not found"
)

var (
	ErrResourceNotFound = errors.New(RESOURCE_NOT_FOUND) // 404
)

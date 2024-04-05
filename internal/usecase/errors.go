package usecase

import (
	"errors"
	"fmt"
)

const LogResource = "study_set"

// ErrNotFound means that resource was not found.
type ErrNotFound struct {
	Resource string
}

func (err *ErrNotFound) Error() string {
	return fmt.Sprintf("resource not found: [%s]", err.Resource)
}

var (
	// ErrDataStoreFailed means that repository failed to complete a required operation.
	ErrDataStoreFailed = errors.New("repository failed")
)

package captures

import (
	"context"
	"errors"
)

type Repository interface {
	// Allocate attempts to create a new entry for Capture. Successful invocation returns Allocation with StatusAllocated.
	// If a Capture has been completed using Record function, Allocation would have StatusCompleted.
	// In all other cases, Allocation has StatusInFlightCapture.
	Allocate(ctx context.Context, idempotencyKey string) (*Allocation, error)

	// Record completes the allocation by persisting Capture.
	// If Allocation does not exist or has been completed previously, it returns ErrConflictOrMissing.
	Record(ctx context.Context, allocation *Allocation) error
}

var ErrAllocationFailed = errors.New("allocation failed")
var ErrRecordFailed = errors.New("record failed")
var ErrConflictOrMissing = errors.New("duplicate or missing entry")

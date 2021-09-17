package captures_mock

import (
	"checkr.com/idempotent-requests/pkg/captures"
	"context"
)

type RepositoryImpl struct {
	Storage map[string]*captures.Allocation
}

func NewRepositoryImpl() *RepositoryImpl {
	storage := make(map[string]*captures.Allocation, 10)

	return &RepositoryImpl{
		Storage: storage,
	}
}

func (r *RepositoryImpl) Allocate(ctx context.Context, idempotencyKey string) (allocation *captures.Allocation, err error) {

	allocation, ok := r.Storage[idempotencyKey]

	if ok {
		switch allocation.Status {
		case captures.StatusCompleted:
			return allocation, nil
		default:
			allocation.Status = captures.StatusInFlightCapture
			return allocation, nil
		}
	} else {
		allocation = &captures.Allocation{
			IdempotencyKey: idempotencyKey,
			Status:         captures.StatusAllocated,
		}
		r.Storage[idempotencyKey] = allocation
	}

	return allocation, nil
}

func (r *RepositoryImpl) Record(ctx context.Context, allocation *captures.Allocation) error {

	allocation, ok := r.Storage[allocation.IdempotencyKey]

	if ok {
		if allocation.Status == captures.StatusAllocated {
			r.Storage[allocation.IdempotencyKey] = allocation
			return nil
		}
	}

	return captures.ErrConflictOrMissing

}

func (r *RepositoryImpl) Ready(ctx context.Context) bool {
	return true
}

package captures_mongo

import (
	"checkr.com/idempotent-requests/pkg/captures"
	"checkr.com/idempotent-requests/pkg/util"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func (repo *RepositoryImpl) Allocate(ctx context.Context, idempotencyKey string) (*captures.Allocation, error) {

	allocation := &captures.Allocation{
		IdempotencyKey: idempotencyKey,
	}

	filter := allocation

	update := Upsert{
		Set:         allocation,
		SetOnInsert: SetStatusOnInsert(captures.StatusAllocated),
	}

	opts := options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.Before)

	result := repo.collection.FindOneAndUpdate(ctx, filter, update, opts)

	if result.Err() != nil {
		// It was the first upsert with a given idempotency key, it means that capture was only allocated.
		if result.Err() == mongo.ErrNoDocuments {
			allocation.Status = captures.StatusAllocated
			return allocation, nil
		} else {
			zap.S().Errorw("failed to update a capture", util.KeyErr, result.Err().Error())
			return allocation, captures.ErrAllocationFailed
		}
	}

	existingAllocation := &captures.Allocation{}
	if err := result.Decode(existingAllocation); err != nil {
		zap.S().Errorw("failed to decode completed capture", util.KeyErr, err.Error())
		return allocation, captures.ErrAllocationFailed
	}

	switch existingAllocation.Status {
	case captures.StatusCompleted:
		return existingAllocation, nil
	default:
		allocation.Status = captures.StatusInFlightCapture
		return allocation, nil
	}

}

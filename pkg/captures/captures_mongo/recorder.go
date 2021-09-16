package captures_mongo

import (
	"checkr.com/idempotent-requests/pkg/captures"
	"checkr.com/idempotent-requests/pkg/util"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func (repo *RepositoryImpl) Record(ctx context.Context, allocation *captures.Allocation) (err error) {

	filter := captures.Allocation{
		IdempotencyKey: allocation.IdempotencyKey,
		Status:         captures.StatusAllocated,
	}

	update := Upsert{
		Set: allocation,
	}

	opts := options.FindOneAndUpdate().
		SetReturnDocument(options.Before)

	result := repo.Collection.FindOneAndUpdate(ctx, filter, update, opts)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			zap.S().Infow("an attempt to update a completed or non-existent capture", util.KeyIdempotencyKey, allocation.IdempotencyKey)
			return captures.ErrConflictOrMissing
		} else {
			zap.S().Errorw("update failed", util.KeyErr, result.Err().Error())
			return captures.ErrRecordFailed
		}
	}

	return nil

}

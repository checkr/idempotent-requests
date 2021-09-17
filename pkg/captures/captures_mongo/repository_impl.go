package captures_mongo

import (
	"checkr.com/idempotent-requests/pkg/captures"
	"checkr.com/idempotent-requests/pkg/mongodb"
	"context"
	"github.com/avast/retry-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"go.uber.org/zap"
	"time"
)

const (
	Database   = "idempotent_requests"
	Collection = "captures"
)

type RepositoryImpl struct {
	client     *mongodb.Client
	collection *mongo.Collection
	ready      bool
}

func NewRepository(client *mongodb.Client) *RepositoryImpl {
	wcMajority := writeconcern.New(writeconcern.WMajority(), writeconcern.WTimeout(30*time.Second))
	wcMajorityCollectionOpts := options.Collection().SetWriteConcern(wcMajority)
	collection := client.Mongo.Database(Database).
		Collection(Collection, wcMajorityCollectionOpts)

	createIndex(context.Background(), collection)

	return &RepositoryImpl{
		client:     client,
		collection: collection,
		ready:      true,
	}
}

func createIndex(ctx context.Context, collection *mongo.Collection) {
	maxRetries := 20
	err := retry.Do(
		func() error {
			indexModel := mongo.IndexModel{
				Keys: captures.AllocationIndex,
				Options: options.Index().
					SetUnique(true).
					SetExpireAfterSeconds(int32(24 * time.Hour.Seconds())),
			}

			if _, err := collection.Indexes().CreateOne(ctx, indexModel); err == nil {
				return nil
			} else {
				return err
			}
		},
		retry.Attempts(uint(maxRetries)),
		retry.MaxDelay(5*time.Second),
		retry.OnRetry(func(n uint, err error) {
			zap.S().Infof("faield to create indexes in mongo. re-try %d/%d", n+1, maxRetries)
		}),
	)

	if err != nil {
		zap.S().Panic(err)
	}
}

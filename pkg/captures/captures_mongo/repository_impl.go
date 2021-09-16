package captures_mongo

import (
	"checkr.com/idempotent-requests/pkg/captures"
	"checkr.com/idempotent-requests/pkg/mongodb"
	"context"
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
	Client     *mongodb.Client
	Collection *mongo.Collection
}

func NewRepository(client *mongodb.Client) *RepositoryImpl {
	wcMajority := writeconcern.New(writeconcern.WMajority(), writeconcern.WTimeout(30*time.Second))
	wcMajorityCollectionOpts := options.Collection().SetWriteConcern(wcMajority)
	collection := client.Mongo.Database(Database).
		Collection(Collection, wcMajorityCollectionOpts)

	createIndex(context.Background(), collection)

	return &RepositoryImpl{
		Client:     client,
		Collection: collection,
	}
}

func createIndex(ctx context.Context, collection *mongo.Collection) {

	indexModel := mongo.IndexModel{
		Keys: captures.AllocationIndex,
		Options: options.Index().
			SetUnique(true).
			SetExpireAfterSeconds(int32(24 * time.Hour.Seconds())),
	}

	if _, err := collection.Indexes().CreateOne(ctx, indexModel); err != nil {
		zap.S().Panic(err)
	}
}

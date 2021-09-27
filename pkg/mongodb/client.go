package mongodb

import (
	"checkr.com/idempotent-requests/pkg/tracing"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"time"
)

type Client struct {
	Mongo *mongo.Client
}

func NewClient(tracer tracing.Tracer) *Client {
	cfg := GetConfig()
	opts := options.Client().ApplyURI(cfg.URI)

	opts = tracer.MongoDB(opts)

	client, err := mongo.NewClient(opts)
	if err != nil {
		zap.S().Panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		zap.S().Panic(err)
	}

	return &Client{Mongo: client}
}

func (c *Client) Shutdown(ctx context.Context) {
	if err := c.Mongo.Disconnect(ctx); err != nil {
		zap.S().Panic(err)
	}
}

package persistence

import (
	"context"
	"fmt"

	"github.com/yourname/classwork/backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB represents a MongoDB connection
type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

// NewMongoDB creates a new MongoDB instance
func NewMongoDB(cfg *config.DatabaseConfig) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ConnTimeout)
	defer cancel()

	opts := options.Client().
		ApplyURI(cfg.URI).
		SetMaxPoolSize(cfg.MaxPoolSize).
		SetMinPoolSize(cfg.MinPoolSize)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	db := client.Database(cfg.Name)

	return &MongoDB{
		client: client,
		db:     db,
	}, nil
}

// GetDB returns the MongoDB database instance
func (m *MongoDB) GetDB() *mongo.Database {
	return m.db
}

// GetClient returns the MongoDB client instance
func (m *MongoDB) GetClient() *mongo.Client {
	return m.client
}

// Close closes the MongoDB connection
func (m *MongoDB) Close(ctx context.Context) error {
	if m.client != nil {
		return m.client.Disconnect(ctx)
	}
	return nil
}

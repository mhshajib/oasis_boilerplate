package seeder

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// Seeder represents a seed contract
type Seeder interface {
	Name() string
	Seed(ctx context.Context, d *mongo.Database) error
}

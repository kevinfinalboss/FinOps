package repository

import (
	"context"

	"github.com/kevinfinalboss/FinOps/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type SpendingRepository struct {
	db *mongo.Collection
}

func NewSpendingRepository(db *mongo.Database, collectionName string) *SpendingRepository {
	return &SpendingRepository{
		db: db.Collection(collectionName),
	}
}

func (r *SpendingRepository) CreateSpending(ctx context.Context, spending domain.Spending) error {
	_, err := r.db.InsertOne(ctx, spending)
	return err
}

package repository

import (
	"context"

	"github.com/kevinfinalboss/FinOps/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type IncomeRepository struct {
	db *mongo.Collection
}

func NewIncomeRepository(db *mongo.Database, collectionName string) *IncomeRepository {
	return &IncomeRepository{
		db: db.Collection(collectionName),
	}
}

func (r *IncomeRepository) CreateIncome(ctx context.Context, income domain.Income) error {
	_, err := r.db.InsertOne(ctx, income)
	return err
}

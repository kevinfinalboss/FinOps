package repository

import (
	"context"

	"github.com/kevinfinalboss/FinOps/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *IncomeRepository) GetRecentIncomes(ctx context.Context) ([]domain.Income, error) {
	var incomes []domain.Income
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}).SetLimit(3)
	cursor, err := r.db.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var income domain.Income
		if err := cursor.Decode(&income); err != nil {
			return nil, err
		}
		incomes = append(incomes, income)
	}
	return incomes, nil
}

func (r *IncomeRepository) GetIncomesSum(ctx context.Context) (float64, error) {
	pipeline := []bson.M{
		{"$group": bson.M{"_id": nil, "total": bson.M{"$sum": "$value"}}},
	}
	cursor, err := r.db.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		return 0, err
	}

	if len(results) == 0 {
		return 0, nil
	}

	total := results[0]["total"].(float64)
	return total, nil
}

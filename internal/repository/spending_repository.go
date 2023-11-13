package repository

import (
	"context"
	"fmt"

	"github.com/kevinfinalboss/FinOps/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *SpendingRepository) GetRecentSpendings(ctx context.Context) ([]domain.Spending, error) {
	var spendings []domain.Spending
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}).SetLimit(3)
	cursor, err := r.db.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var spending domain.Spending
		if err := cursor.Decode(&spending); err != nil {
			return nil, err
		}
		spendings = append(spendings, spending)
	}
	return spendings, nil
}

func (r *SpendingRepository) GetSpendingsByMonth(ctx context.Context, month string) ([]domain.Spending, error) {
	var spendings []domain.Spending
	regexPattern := fmt.Sprintf("^.*%s.*$", month)
	query := bson.M{"date": bson.M{"$regex": regexPattern}}
	cursor, err := r.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var spending domain.Spending
		if err := cursor.Decode(&spending); err != nil {
			return nil, err
		}
		spendings = append(spendings, spending)
	}
	return spendings, nil
}

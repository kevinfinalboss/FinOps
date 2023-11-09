package repository

import (
	"context"

	"github.com/kevinfinalboss/FinOps/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collectionName string) *UserRepository {
	return &UserRepository{
		db: db.Collection(collectionName),
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user domain.User) error {
	_, err := r.db.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := r.db.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

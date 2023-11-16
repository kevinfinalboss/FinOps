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
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) IsEmailInUse(ctx context.Context, email string) bool {
	var user domain.User
	err := r.db.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return err != mongo.ErrNoDocuments
}

func (r *UserRepository) FindUserByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	err := r.db.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	var users []*domain.User
	cursor, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) SaveRefreshToken(ctx context.Context, userID string, refreshToken string) error {
	_, err := r.db.UpdateOne(
		ctx,
		bson.M{"_id": userID},
		bson.M{"$set": bson.M{"refresh_token": refreshToken}},
	)
	return err
}

func (r *UserRepository) RemoveRefreshToken(ctx context.Context, userID string) error {
	_, err := r.db.UpdateOne(
		ctx,
		bson.M{"_id": userID},
		bson.M{"$unset": bson.M{"refresh_token": ""}},
	)
	return err
}

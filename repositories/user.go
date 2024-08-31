package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/fleimkeipa/challengers-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (rc *UserRepository) Create(ctx context.Context, user model.User) (model.User, error) {
	hashedPassword, err := model.HashPassword(user.Password)
	if err != nil {
		return model.User{}, err
	}
	user.Password = hashedPassword

	query, err := rc.
		db.
		Collection("users").
		InsertOne(ctx, &user)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	oid, ok := query.InsertedID.(primitive.ObjectID)
	if !ok {
		return model.User{}, errors.New("can't get inserted ID")
	}

	user.ID = oid.Hex()

	return user, nil
}

func (rc *UserRepository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	var user = new(model.User)
	err := rc.
		db.
		Collection("users").
		FindOne(ctx, bson.M{"username": username}).
		Decode(user)
	if err != nil {
		return model.User{}, err
	}

	return *user, nil
}

package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/fleimkeipa/challengers-api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection = "users"

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (rc *UserRepository) Create(ctx context.Context, user model.User) (model.User, error) {
	query, err := rc.
		db.
		Collection(userCollection).
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
		Collection(userCollection).
		FindOne(ctx, bson.M{"username": username}).
		Decode(user)
	if err != nil {
		return model.User{}, err
	}

	return *user, nil
}

func (rc *UserRepository) Get(ctx context.Context, opts model.UserFindOpts) ([]model.User, error) {
	findOpts, filter := userFilters(opts)

	cur, err := rc.
		db.
		Collection(userCollection).
		Find(ctx, filter, &findOpts)
	if err != nil {
		return []model.User{}, err
	}

	var users = make([]model.User, 0)
	if err := cur.All(ctx, &users); err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func userFilters(opts model.UserFindOpts) (options.FindOptions, bson.M) {
	var findOpts = getPaginationOpts(opts.Limit, opts.Skip)

	var filter = bson.M{}
	switch {
	case opts.Username.IsActive:
		filter = bson.M{"username": opts.Username.Value}
	case opts.RoleID.IsActive:
		filter = bson.M{"role_id": opts.RoleID.Value}
	case opts.Email.IsActive:
		filter = bson.M{"email": opts.Email.Value}
	}

	return findOpts, filter
}

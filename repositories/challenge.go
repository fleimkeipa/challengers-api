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

var chCollection = "challenges"

type ChallengeRepository struct {
	db *mongo.Database
}

func NewChallengeRepository(db *mongo.Database) *ChallengeRepository {
	return &ChallengeRepository{
		db: db,
	}
}

func (rc *ChallengeRepository) Create(ctx context.Context, challenge model.Challenge) (model.Challenge, error) {
	query, err := rc.
		db.
		Collection(chCollection).
		InsertOne(ctx, &challenge)
	if err != nil {
		return model.Challenge{}, fmt.Errorf("failed to create user: %w", err)
	}

	oid, ok := query.InsertedID.(primitive.ObjectID)
	if !ok {
		return model.Challenge{}, errors.New("can't get inserted ID")
	}

	challenge.ID = oid.Hex()

	return challenge, nil
}

func (rc *ChallengeRepository) Update(ctx context.Context, challenge model.Challenge) (model.Challenge, error) {
	oID, err := primitive.ObjectIDFromHex(challenge.ID)
	if err != nil {
		return model.Challenge{}, fmt.Errorf("failed to convert id: %w", err)
	}

	query, err := rc.
		db.
		Collection(chCollection).
		UpdateOne(ctx, &challenge, bson.M{"_id": oID})
	if err != nil {
		return model.Challenge{}, fmt.Errorf("failed to create user: %w", err)
	}

	if query.MatchedCount == 0 {
		return model.Challenge{}, fmt.Errorf("not found challenge with id: %v", challenge.ID)
	}

	return challenge, nil
}

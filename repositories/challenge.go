package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

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

	var filter = bson.M{"_id": oID}
	query, err := rc.
		db.
		Collection(chCollection).
		UpdateOne(ctx, &challenge, filter)
	if err != nil {
		return model.Challenge{}, fmt.Errorf("failed to create user: %w", err)
	}

	if query.MatchedCount == 0 {
		return model.Challenge{}, fmt.Errorf("not found challenge with id: %v", challenge.ID)
	}

	return challenge, nil
}

func (rc *ChallengeRepository) Delete(ctx context.Context, id string) error {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert id: %w", err)
	}

	var filter = bson.M{"_id": oID}
	var updater = bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
			"is_active":  0,
		},
	}
	query, err := rc.
		db.
		Collection(chCollection).
		UpdateOne(ctx, filter, updater)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	if query.MatchedCount == 0 {
		return fmt.Errorf("not found challenge with id: %v", id)
	}

	return nil
}

func (rc *ChallengeRepository) GetByID(ctx context.Context, id string) (model.Challenge, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Challenge{}, fmt.Errorf("failed to convert id: %w", err)
	}

	var challenge = new(model.Challenge)
	err = rc.
		db.
		Collection(chCollection).
		FindOne(ctx, bson.M{"_id": oID}).
		Decode(challenge)
	if err != nil {
		return model.Challenge{}, err
	}

	return *challenge, nil
}

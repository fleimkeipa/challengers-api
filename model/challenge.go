package model

import "time"

type ChallengeRequest struct {
	ID   string `bson:"_id,omitempty"` // MongoDB ObjectID, primary key
	Name string `bson:"name" json:"name"`
}

type Challenge struct {
	ID        string    `bson:"_id,omitempty"` // MongoDB ObjectID, primary key
	Name      string    `bson:"name" json:"name"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

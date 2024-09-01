package model

import "time"

type ChallengeRequest struct {
	ID   string `bson:"_id,omitempty"` // MongoDB ObjectID, primary key
	Name string `bson:"name" json:"name"`
}

type Challenge struct {
	ID        string    `bson:"_id,omitempty"` // MongoDB ObjectID, primary key
	Name      string    `bson:"name" json:"name"`
	IsActive  bool      `bson:"is_active" json:"is_active"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at"`
	DeletedAt time.Time `bson:"deleted_at,omitempty" json:"deleted_at"`
}

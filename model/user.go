package model

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	ID       string `bson:"_id,omitempty"` // MongoDB ObjectID, primary key
	RoleID   uint   `bson:"role_id" json:"role_id"`
	Username string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"-"`
}

// VerifyPassword verifies if the given password matches the stored hash.
func ValidateUserPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// HashPassword generates a bcrypt hash for the given password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(bytes), nil
}

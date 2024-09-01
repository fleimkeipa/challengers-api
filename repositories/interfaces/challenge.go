package interfaces

import (
	"context"

	"github.com/fleimkeipa/challengers-api/model"
)

type ChallengeInterfaces interface {
	Create(context.Context, model.Challenge) (model.Challenge, error)
	Update(context.Context, model.Challenge) (model.Challenge, error)
	Delete(context.Context, string) error
	Get(context.Context, model.ChallengeFindOpts) ([]model.Challenge, error)
	GetByID(context.Context, string) (model.Challenge, error)
}

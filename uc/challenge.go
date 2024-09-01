package uc

import (
	"context"
	"time"

	"github.com/fleimkeipa/challengers-api/model"
	"github.com/fleimkeipa/challengers-api/repositories/interfaces"
)

type ChallengeUC struct {
	repo interfaces.ChallengeInterfaces
}

func NewChallengeUC(repo interfaces.ChallengeInterfaces) *ChallengeUC {
	return &ChallengeUC{
		repo: repo,
	}
}

func (rc *ChallengeUC) Create(ctx context.Context, challenge model.Challenge) (model.Challenge, error) {
	challenge.CreatedAt = time.Now()

	return rc.repo.Create(ctx, challenge)
}

func (rc *ChallengeUC) Update(ctx context.Context, challenge model.Challenge) (model.Challenge, error) {
	challenge.UpdatedAt = time.Now()

	return rc.repo.Update(ctx, challenge)
}

package uc

import (
	"context"
	"time"

	"github.com/fleimkeipa/challengers-api/model"
	"github.com/fleimkeipa/challengers-api/repositories/interfaces"
)

type ChallengeUC struct {
	chRepo interfaces.ChallengeInterfaces
}

func NewChallengeUC(chRepo interfaces.ChallengeInterfaces) *ChallengeUC {
	return &ChallengeUC{
		chRepo: chRepo,
	}
}

func (rc *ChallengeUC) Create(ctx context.Context, challenge model.Challenge) (model.Challenge, error) {
	challenge.CreatedAt = time.Now()

	return rc.chRepo.Create(ctx, challenge)
}

func (rc *ChallengeUC) Update(ctx context.Context, challenge model.Challenge) (model.Challenge, error) {
	challenge.UpdatedAt = time.Now()

	return rc.chRepo.Update(ctx, challenge)
}

func (rc *ChallengeUC) Delete(ctx context.Context, id string) error {
	return rc.chRepo.Delete(ctx, id)
}

func (rc *ChallengeUC) Get(ctx context.Context, opts model.ChallengeFindOpts) ([]model.Challenge, error) {
	return rc.chRepo.Get(ctx, opts)
}

func (rc *ChallengeUC) GetByID(ctx context.Context, id string) (model.Challenge, error) {
	return rc.chRepo.GetByID(ctx, id)
}

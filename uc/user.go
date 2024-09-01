package uc

import (
	"context"

	"github.com/fleimkeipa/challengers-api/model"
	"github.com/fleimkeipa/challengers-api/repositories/interfaces"
)

type UserUC struct {
	repo interfaces.UserInterfaces
}

func NewUserUC(repo interfaces.UserInterfaces) *UserUC {
	return &UserUC{
		repo: repo,
	}
}

func (rc *UserUC) Create(ctx context.Context, user model.User) (model.User, error) {
	hashedPassword, err := model.HashPassword(user.Password)
	if err != nil {
		return model.User{}, err
	}
	user.Password = hashedPassword

	return rc.repo.Create(ctx, user)
}

func (rc *UserUC) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	return rc.repo.GetUserByUsername(ctx, username)
}

func (rc *UserUC) Get(ctx context.Context, opts model.UserFindOpts) ([]model.User, error) {
	return rc.repo.Get(ctx, opts)
}

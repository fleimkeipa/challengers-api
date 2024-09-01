package interfaces

import (
	"context"

	"github.com/fleimkeipa/challengers-api/model"
)

type UserInterfaces interface {
	Create(context.Context, model.User) (model.User, error)
	Get(context.Context, model.UserFindOpts) ([]model.User, error)
	GetUserByUsername(context.Context, string) (model.User, error)
}

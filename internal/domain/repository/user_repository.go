package repository

import (
	"context"

	"github.com/ei-sugimoto/techGO/internal/domain/model"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, userId string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
}

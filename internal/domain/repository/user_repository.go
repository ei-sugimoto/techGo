package repository

import (
	"context"

	"github.com/ei-sugimoto/techGO/internal/domain/model"
)


type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
}

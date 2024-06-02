package usecase

import (
	"context"

	"github.com/ei-sugimoto/techGO/internal/domain/service"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, name string) (context.Context, error)
}

type userUsecase struct {
	userService *service.UserService
}

func NewUserUsecase(userService *service.UserService) UserUsecase {
	return &userUsecase{
		userService: userService,
	}
}

func (u *userUsecase) CreateUser(ctx context.Context, name string) (context.Context, error) {
	return u.userService.CreateUser(ctx, name)
}

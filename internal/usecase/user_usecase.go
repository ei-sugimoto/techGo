package usecase

import (
	"context"

	"github.com/ei-sugimoto/techGO/internal/domain/service"
	input "github.com/ei-sugimoto/techGO/internal/usecase/Input"
)

type IUserUseCase interface {
	CreateUser(ctx context.Context, i input.UserInput) (context.Context, error)
}

type userUseCase struct {
	userService *service.UserService
}

func NewUserUsecase(userService *service.UserService) IUserUseCase {
	return &userUseCase{
		userService: userService,
	}
}

func (u *userUseCase) CreateUser(ctx context.Context, i input.UserInput) (context.Context, error) {
	return u.userService.CreateUser(ctx, i.Name)
}

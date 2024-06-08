package usecase

import (
	"context"

	"github.com/ei-sugimoto/techGO/internal/domain/service"
	input "github.com/ei-sugimoto/techGO/internal/usecase/Input"
	"github.com/ei-sugimoto/techGO/internal/usecase/output"
)

type IUserUseCase interface {
	CreateUser(ctx context.Context, i input.CreateUserInput) (context.Context, error)
	GetUser(ctx context.Context, i *input.GetUserInput) (context.Context, *output.GetUserOutput, error)
	UpdateUser(ctx context.Context, i *input.UpdateUserInput) (context.Context, error)
}

type userUseCase struct {
	userService *service.UserService
}

func NewUserUsecase(userService *service.UserService) IUserUseCase {
	return &userUseCase{
		userService: userService,
	}
}

func (u *userUseCase) CreateUser(ctx context.Context, i input.CreateUserInput) (context.Context, error) {
	return u.userService.CreateUser(ctx, i.Name)
}

func (u *userUseCase) GetUser(ctx context.Context, i *input.GetUserInput) (context.Context, *output.GetUserOutput, error) {
	ctx, user, err := u.userService.GetUser(ctx, i.UserID)
	if err != nil {
		return ctx, nil, err
	}
	return ctx, &output.GetUserOutput{
		UserID: user.UserID.String(),
		Name:   user.Name,
	}, nil
}

func (u *userUseCase) UpdateUser(ctx context.Context, i *input.UpdateUserInput) (context.Context, error) {
	return u.userService.UpdateUser(ctx, i.UserID, i.Name)
}

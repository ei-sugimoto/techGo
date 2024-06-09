package usecase

import (
	"context"

	"github.com/ei-sugimoto/techGO/internal/domain/model"
	"github.com/ei-sugimoto/techGO/internal/domain/service"
	input "github.com/ei-sugimoto/techGO/internal/usecase/Input"
)

type IUserCharacterUseCase interface {
	GetUserCharacter(ctx context.Context, i *input.GetUserCharacterInput) (context.Context, []*model.UserCharacter, error)
}

type UserCharacterUseCase struct {
	userCharacterSerivce *service.UserCharacterService
}

func NewUserCharacterUseCase(userCharacterSerivce *service.UserCharacterService) IUserCharacterUseCase {
	return &UserCharacterUseCase{
		userCharacterSerivce: userCharacterSerivce,
	}
}

func (u *UserCharacterUseCase) GetUserCharacter(ctx context.Context, i *input.GetUserCharacterInput) (context.Context, []*model.UserCharacter, error) {
	newctx, rows, err := u.userCharacterSerivce.GetUserCharacter(ctx, i.UserID)

	if err != nil {
		return newctx, nil, err
	}
	return newctx, rows, nil
}

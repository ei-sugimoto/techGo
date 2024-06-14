package usecase

import (
	"context"

	"github.com/ei-sugimoto/techGO/internal/domain/service"
	input "github.com/ei-sugimoto/techGO/internal/usecase/Input"
	"github.com/ei-sugimoto/techGO/internal/usecase/output"
)

type IUserCharacterUseCase interface {
	GetUserCharacter(ctx context.Context, i *input.GetUserCharacterInput) (context.Context, output.GetUserCharacterOutputs, error)
	CreateUserCharacter(ctx context.Context, i *input.CreateUserCharacterInput) (context.Context, output.CreateUserCharacterOutputs, error)
}

type UserCharacterUseCase struct {
	userCharacterSerivce *service.UserCharacterService
}

func NewUserCharacterUseCase(userCharacterSerivce *service.UserCharacterService) IUserCharacterUseCase {
	return &UserCharacterUseCase{
		userCharacterSerivce: userCharacterSerivce,
	}
}

func (u *UserCharacterUseCase) GetUserCharacter(ctx context.Context, i *input.GetUserCharacterInput) (context.Context,
	output.GetUserCharacterOutputs, error) {
	newctx, rows, err := u.userCharacterSerivce.GetUserCharacter(ctx, i.UserID)

	if err != nil {
		return newctx, nil, err
	}
	res := output.GetUserCharacterOutputs{}
	for _, row := range rows {
		res = append(res, output.GetUserCharacterOutput{
			UserCharacterID: row.UserCharacterID,
			Name:            row.Character.Name,
			CharacterID:     row.Character.CharacterID,
		})
	}
	return newctx, res, nil
}

func (u *UserCharacterUseCase) CreateUserCharacter(ctx context.Context, i *input.CreateUserCharacterInput) (context.Context,
	output.CreateUserCharacterOutputs, error) {
	newctx, rows, err := u.userCharacterSerivce.CreateUserCharacter(ctx, i.UserID, i.Times)

	if err != nil {
		return newctx, nil, err
	}
	res := output.CreateUserCharacterOutputs{}
	for _, row := range rows {
		res = append(res, output.CreateUserCharacterOutput{
			Name:        row.Name,
			CharacterID: row.CharacterID,
		})
	}
	return newctx, res, nil
}

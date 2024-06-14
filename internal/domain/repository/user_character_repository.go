package repository

import (
	"context"

	"github.com/ei-sugimoto/techGO/internal/domain/model"
)

type IUserCharacterRepository interface {
	GetUserChraracter(ctx context.Context, userId string) ([]*model.UserCharacter, error)
	CreateUserCharacter(ctx context.Context, userId string, times int) ([]*model.Character, error)
}

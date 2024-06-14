package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ei-sugimoto/techGO/internal/domain/model"
	"github.com/ei-sugimoto/techGO/internal/domain/repository"
	"github.com/ei-sugimoto/techGO/pkg"
)

type UserCharacterService struct {
	userCharacterRepository repository.IUserCharacterRepository
	logger                  *slog.Logger
}

func NewUserCharacterService(userCharacterRepository repository.IUserCharacterRepository) *UserCharacterService {
	return &UserCharacterService{userCharacterRepository: userCharacterRepository, logger: pkg.NewLogger().With(slog.String("path", "service/user_character_service.go"))}
}

func (s *UserCharacterService) GetUserCharacter(ctx context.Context, userId string) (context.Context, []*model.UserCharacter, error) {
	userCharacter, err := s.userCharacterRepository.GetUserChraracter(ctx, userId)
	if err != nil {
		s.logger.Error(err.Error())
		return ctx, nil, err
	}
	return ctx, userCharacter, nil
}

func (s *UserCharacterService) CreateUserCharacter(ctx context.Context, userId string, times int) (context.Context, []*model.Character, error) {
	characters, err := s.userCharacterRepository.CreateUserCharacter(ctx, userId, times)
	if err != nil {
		s.logger.Error(err.Error())
		return ctx, nil, err
	}
	s.logger.Info(fmt.Sprintf("created user characters: %v", characters))
	return ctx, characters, nil
}

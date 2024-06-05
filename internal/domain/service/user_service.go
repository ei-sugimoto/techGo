package service

import (
	"context"
	"log/slog"

	"github.com/ei-sugimoto/techGO/internal/domain/model"
	"github.com/ei-sugimoto/techGO/internal/domain/repository"
	"github.com/ei-sugimoto/techGO/pkg"
	"github.com/google/uuid"
)

type UserService struct {
	userRepository repository.IUserRepository
	logger         *slog.Logger
}

func NewUserService(userRepository repository.IUserRepository) *UserService {
	return &UserService{userRepository: userRepository, logger: pkg.NewLogger().With(slog.String("path", "service/user_service.go"))}
}

func (s *UserService) CreateUser(ctx context.Context, name string) (context.Context, error) {

	uuid := uuid.New()
	user := &model.User{
		UserID: uuid,
		Name:   name,
	}
	err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		s.logger.Error(err.Error())
		return ctx, err
	}
	token := pkg.EncodeJwt(uuid.String())
	s.logger.Info("created Token", slog.String("token", token))
	ctxWithToken := context.WithValue(context.Background(), "token", token)
	return ctxWithToken, nil
}

func (s *UserService) GetUser(ctx context.Context, userId string) (context.Context, *model.User, error) {
	user, err := s.userRepository.GetUser(ctx, userId)
	if err != nil {
		s.logger.Error(err.Error())
		return ctx, nil, err
	}
	return ctx, user, nil
}

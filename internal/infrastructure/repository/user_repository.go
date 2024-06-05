package repository

import (
	"context"

	"github.com/ei-sugimoto/techGO/internal/domain/model"
	"github.com/ei-sugimoto/techGO/internal/domain/repository"
	"github.com/ei-sugimoto/techGO/internal/infrastructure/dao"
)

type userRepository struct {
	DB *dao.DataBase
}

func NewUserRepository(db *dao.DataBase) repository.IUserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	return r.DB.GormDB.Create(user).Error
}

func (r *userRepository) GetUser(ctx context.Context, userId string) (*model.User, error) {
	user := &model.User{}
	if err := r.DB.GormDB.Where("user_id = ?", userId).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *model.User) error {
	// ここでユーザーの更新処理を実装します
	return nil
}

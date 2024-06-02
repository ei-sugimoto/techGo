package repository

import (
	"context"

	"github.com/ei-sugimoto/techGO/internal/domain/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {

	return r.db.Create(user).Error
}

func (r *userRepository) GetUser(ctx context.Context, user *model.User) error {
	// ここでユーザーの取得処理を実装します
	return nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *model.User) error {
	// ここでユーザーの更新処理を実装します
	return nil
}

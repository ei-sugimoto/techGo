package repository

import (
	"context"
	"log/slog"

	"github.com/ei-sugimoto/techGO/internal/domain/model"
	"github.com/ei-sugimoto/techGO/internal/domain/repository"
	"github.com/ei-sugimoto/techGO/internal/infrastructure/dao"
	"github.com/ei-sugimoto/techGO/pkg"
)

type UserCharacterRepository struct {
	DB     *dao.DataBase
	logger *slog.Logger
}

func NewUserCharacterRepository(db *dao.DataBase) repository.IUserCharacterRepository {
	logger := pkg.NewLogger()
	return &UserCharacterRepository{DB: db, logger: logger}
}

func (r *UserCharacterRepository) GetUserChraracter(ctx context.Context, userId string) ([]*model.UserCharacter, error) {
	rows := []*model.UserCharacter{}

	if err := r.DB.GormDB.Model(&model.UserCharacter{}).Where("user_id = ?", userId).Find(&rows).Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	if len(rows) == 0 {
		r.logger.Error("user_character not found")
		return nil, nil
	}
	return rows, nil
}

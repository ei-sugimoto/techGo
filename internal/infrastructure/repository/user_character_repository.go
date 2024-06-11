package repository

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/ei-sugimoto/techGO/internal/domain/model"
	"github.com/ei-sugimoto/techGO/internal/domain/repository"
	"github.com/ei-sugimoto/techGO/internal/infrastructure/dao"
	"github.com/ei-sugimoto/techGO/pkg"
	llog "gorm.io/gorm/logger"
)

type UserCharacterRepository struct {
	DB     *dao.DataBase
	logger *slog.Logger
}

func NewUserCharacterRepository(db *dao.DataBase) repository.IUserCharacterRepository {
	logger := pkg.NewLogger()
	newLogger := llog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		llog.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      llog.Info,   // Log level
			Colorful:      false,       // Disable color
		},
	)

	db.GormDB.Config.Logger = newLogger
	return &UserCharacterRepository{DB: db, logger: logger}
}

func (r *UserCharacterRepository) GetUserChraracter(ctx context.Context, userId string) ([]*model.UserCharacter, error) {
	var userCharacters []*model.UserCharacter

	if err := r.DB.GormDB.Where("user_character.user_id = ?", userId).
		Joins("User").
		Joins("Character").
		Preload("User").
		Preload("Character").
		Find(&userCharacters).Error; err != nil {
		r.logger.Error(fmt.Sprintf("failed to get user characters: %v", err))
		return nil, err
	}

	fmt.Println(userCharacters[0])
	return userCharacters, nil
}

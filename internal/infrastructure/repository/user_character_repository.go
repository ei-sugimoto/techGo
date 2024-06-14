package repository

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"math/rand"
	"os"
	"time"

	"github.com/ei-sugimoto/techGO/internal/domain/model"
	"github.com/ei-sugimoto/techGO/internal/domain/repository"
	"github.com/ei-sugimoto/techGO/internal/infrastructure/dao"
	"github.com/ei-sugimoto/techGO/pkg"
	"github.com/google/uuid"
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

	if err := r.DB.GormDB.Where("user_character.user_id = ?", userId).Model(&model.UserCharacter{}).
		Preload("Character").
		Find(&userCharacters).Error; err != nil {
		r.logger.Error(fmt.Sprintf("failed to get user characters: %v", err))
		return nil, err
	}
	return userCharacters, nil
}

func (r *UserCharacterRepository) CreateUserCharacter(ctx context.Context, userId string, times int) ([]*model.Character, error) {
	var characters []*model.Character
	for len(characters) < times {
		var count int64
		if err := r.DB.GormDB.Model(&model.Character{}).Count(&count).Error; err != nil {
			r.logger.Info(fmt.Sprintf("failed to count characters: %v", err))
			return nil, err
		}

		offset := rand.Intn(int(count))
		character := &model.Character{}
		if err := r.DB.GormDB.Offset(offset).Limit(1).Find(character).Error; err != nil {
			r.logger.Info(fmt.Sprintf("failed to get character: %v", err))
			return nil, err
		}

		characters = append(characters, character)
	}
	r.logger.Info(fmt.Sprintf("characters: %v", characters))
	for _, character := range characters {
		userCharacter := &model.UserCharacter{
			UserID:          userId,
			CharacterID:     character.CharacterID,
			UserCharacterID: uuid.New().String(),
		}
		if err := r.DB.GormDB.Create(userCharacter).Error; err != nil {
			r.logger.Error(fmt.Sprintf("failed to create user character: %v", err))
			return nil, err
		}

	}
	return characters, nil
}

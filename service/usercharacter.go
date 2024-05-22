package service

import (
	"context"
	"database/sql"

	"github.com/ei-sugimoto/techGO/model"
)

// A TODOService implements CRUD of TODO entities.
type UserCharacter struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewUserCharacter(db *sql.DB) *UserCharacter {
	return &UserCharacter{
		db: db,
	}
}

func(s *UserCharacter) GetUserCharacters(ctx context.Context) ([]*model.UserCharacter, error) {
	rows, err := s.db.Query("SELECT * FROM user_character")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var UserCharacters []*model.UserCharacter
	for rows.Next() {
		var userCharacter model.UserCharacter
		if err := rows.Scan(&userCharacter.UserCharacterID, &userCharacter.CharacterID, &userCharacter.Name); err != nil {
			return nil, err
		}
		UserCharacters = append(UserCharacters, &userCharacter)
	}

	return UserCharacters, nil
}
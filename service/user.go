package service

import (
	"context"
	"database/sql"

	"github.com/ei-sugimoto/techGO/model"
)

// A TODOService implements CRUD of TODO entities.
type UserService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func(s *UserService) GetUsers(ctx context.Context) ([]*model.User, error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.UserName, &user.Password, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}
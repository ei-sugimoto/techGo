package service

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/ei-sugimoto/techGO/model"
	"github.com/google/uuid"
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

func(s *UserCharacter) CreateUser(ctx context.Context, req *model.UserCreateRequest) (*model.UserCreateResponse, *model.CutomError) {
	name := req.Name
	if name == ""{
		err := model.CutomError{Code: http.StatusBadRequest, Message: "name is required"}
		log.Println(err.Message)
		return nil, &err 
	}
	uuid := uuid.New().String()

	_, err := s.db.ExecContext(ctx, "INSERT INTO user_character (user_character_id, name) VALUES (?, ?)", uuid, name)
	if err != nil {
		log.Println(err.Error())
		return nil, &model.CutomError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	claims := jwt.MapClaims{
		"user_character_id": uuid,
		"name": name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Println(err.Error())
		return nil, &model.CutomError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return &model.UserCreateResponse{Token: tokenString}, nil
}

func(s *UserCharacter) GetUserCharacter(ctx context.Context, name string, UserCharacterID string)(*model.UserGetResponse, *model.CutomError){
	res, err := s.db.QueryContext(ctx, "SELECT * FROM user_character WHERE user_character_id = ?", name)
	if err != nil {
		return nil, &model.CutomError{Code: http.StatusNotFound, Message: "user not found"}
	}
	userCharacter := model.UserCharacter{}

	if err := res.Scan(&userCharacter.UserCharacterID, &userCharacter.CharacterID, &userCharacter.Name); err != nil {
		return nil, &model.CutomError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	return &model.UserGetResponse{Name: userCharacter.Name}, nil
}
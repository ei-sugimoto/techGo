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
type User struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewUser(db *sql.DB) *User {
	return &User{
		db: db,
	}
}

func (s *User) GetUsers(ctx context.Context) ([]*model.User, error) {
	rows, err := s.db.Query("SELECT * FROM user_character")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Users []*model.User
	for rows.Next() {
		var User model.User
		if err := rows.Scan(&User.UserID, &User.Name); err != nil {
			return nil, err
		}
		Users = append(Users, &User)
	}

	return Users, nil
}

func (s *User) CreateUser(ctx context.Context, req *model.UserCreateRequest) (*model.UserCreateResponse, *model.CutomError) {
	name := req.Name
	if name == "" {
		err := model.CutomError{Code: http.StatusBadRequest, Message: "name is required"}
		log.Println(err.Message)
		return nil, &err
	}
	uuid := uuid.New().String()

	_, err := s.db.ExecContext(ctx, "INSERT INTO user (user_id, name) VALUES (?, ?)", uuid, name)
	if err != nil {
		log.Println(err.Error())
		return nil, &model.CutomError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	claims := jwt.MapClaims{
		"name": name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Println(err.Error())
		return nil, &model.CutomError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	log.Println("User Create tokenString: ", tokenString)
	return &model.UserCreateResponse{Token: tokenString}, nil
}

func (s *User) GetUser(ctx context.Context, token string) (*model.UserGetResponse, *model.CutomError) {
	name, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		log.Println(err.Error())
		return nil, &model.CutomError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	res, err := s.db.QueryContext(ctx, "SELECT * FROM user WHERE user_id = ?", name)
	if err != nil {
		return nil, &model.CutomError{Code: http.StatusNotFound, Message: "user not found"}
	}
	User := model.User{}

	if err := res.Scan(&User.UserID, &User.Name); err != nil {
		return nil, &model.CutomError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	return &model.UserGetResponse{Name: User.Name}, nil
}

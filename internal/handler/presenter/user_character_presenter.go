package presenter

import (
	"context"

	"github.com/ei-sugimoto/techGO/pkg"
	"github.com/gin-gonic/gin"
)

type UserCharacterPresenter struct{}

func NewUserCharacterPresenter() *UserCharacterPresenter {
	return &UserCharacterPresenter{}
}

type UserCharacterGetRequest struct {
	UserID string `json:"user_id"`
}

type UserCharacterGetResponses struct {
	UserCharacters []UserCharacter `json:"user_characters"`
}

type UserCharacter struct {
	UserCharacterID string `json:"user_character_id"`
	CharacterID     string `json:"character_id"`
	Name            string `json:"name"`
}

type UserCharacterCreateRequest struct {
	UserID string `json:"user_id"`
	Times  int    `json:"times"`
}

type UserCharacterCreateResponses struct {
	UserCharacters []UserCharacter `json:"user_characters"`
}

func (p *UserCharacterPresenter) GetUserCharacterRequest(ctx *gin.Context) (*UserCharacterGetRequest, error) {
	var req UserCharacterGetRequest
	token := ctx.GetHeader("x-token")
	userId, err := pkg.DecodeJwt(token, "user_id")
	if err != nil {
		return nil, err
	}
	req.UserID = userId

	return &req, nil
}

func (p *UserCharacterPresenter) GetUserCharacterResponse(ctx context.Context, rows *[]UserCharacter) *UserCharacterGetResponses {
	res := UserCharacterGetResponses{}
	for _, row := range *rows {
		res.UserCharacters = append(res.UserCharacters, UserCharacter{
			UserCharacterID: row.UserCharacterID,
			CharacterID:     row.CharacterID,
			Name:            row.Name,
		})
	}
	return &res
}

func (p *UserCharacterPresenter) CreateUserCharacterRequest(ctx *gin.Context) (*UserCharacterCreateRequest, error) {
	var req UserCharacterCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	token := ctx.GetHeader("x-token")
	userId, err := pkg.DecodeJwt(token, "user_id")
	if err != nil {
		return nil, err
	}
	req.UserID = userId

	return &req, nil
}

func (p *UserCharacterPresenter) CreateUserCharacterResponse(ctx context.Context, rows *[]UserCharacter) *UserCharacterCreateResponses {
	res := UserCharacterCreateResponses{}
	for _, row := range *rows {
		res.UserCharacters = append(res.UserCharacters, UserCharacter{
			CharacterID: row.CharacterID,
			Name:        row.Name,
		})
	}
	return &res
}

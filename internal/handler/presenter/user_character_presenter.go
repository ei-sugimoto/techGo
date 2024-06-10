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

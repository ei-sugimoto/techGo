package handler

import (
	"github.com/ei-sugimoto/techGO/internal/handler/presenter"
	"github.com/ei-sugimoto/techGO/internal/usecase"
	input "github.com/ei-sugimoto/techGO/internal/usecase/Input"
	"github.com/gin-gonic/gin"
)

type IUserCharacterHandler interface {
	GetUserCharacter(c *gin.Context) (*presenter.UserCharacterGetResponses, error)
	CreateUserChaaracter(ctx *gin.Context) (*presenter.UserCharacterCreateResponses, error)
}

type userCharacterHandler struct {
	userCharacterUseCase   usecase.IUserCharacterUseCase
	userCharacterPresenter presenter.UserCharacterPresenter
}

func NewUserCharacterHandler(userCharacterUseCase usecase.IUserCharacterUseCase, userCharacterPresenter presenter.UserCharacterPresenter) IUserCharacterHandler {
	return &userCharacterHandler{userCharacterUseCase: userCharacterUseCase, userCharacterPresenter: userCharacterPresenter}
}

func (h *userCharacterHandler) GetUserCharacter(ctx *gin.Context) (*presenter.UserCharacterGetResponses, error) {
	req, err := h.userCharacterPresenter.GetUserCharacterRequest(ctx)
	if err != nil || req == nil {
		return nil, err
	}
	input := input.GetUserCharacterInput{UserID: req.UserID}
	newCtx, output, err := h.userCharacterUseCase.GetUserCharacter(ctx.Request.Context(), &input)
	if err != nil {
		return nil, err
	}
	var rows []presenter.UserCharacter
	for _, row := range output {
		rows = append(rows, presenter.UserCharacter{
			UserCharacterID: row.UserCharacterID,
			CharacterID:     row.CharacterID,
			Name:            row.Name,
		})
	}

	return h.userCharacterPresenter.GetUserCharacterResponse(newCtx, &rows), nil
}

func (h *userCharacterHandler) CreateUserChaaracter(ctx *gin.Context) (*presenter.UserCharacterCreateResponses, error) {
	req, err := h.userCharacterPresenter.CreateUserCharacterRequest(ctx)
	if err != nil || req == nil {
		return nil, err
	}
	input := input.CreateUserCharacterInput{
		UserID: req.UserID,
		Times:  req.Times,
	}
	newCtx, output, err := h.userCharacterUseCase.CreateUserCharacter(ctx.Request.Context(), &input)
	if err != nil {
		return nil, err
	}
	var rows []presenter.UserCharacter
	for _, row := range output {
		rows = append(rows, presenter.UserCharacter{
			CharacterID: row.CharacterID,
			Name:        row.Name,
		})
	}
	return h.userCharacterPresenter.CreateUserCharacterResponse(newCtx, &rows), nil
}

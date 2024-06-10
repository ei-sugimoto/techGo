package handler

import (
	"github.com/ei-sugimoto/techGO/internal/handler/presenter"
	"github.com/ei-sugimoto/techGO/internal/usecase"
	input "github.com/ei-sugimoto/techGO/internal/usecase/Input"
	"github.com/ei-sugimoto/techGO/pkg"
	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	CreateUser(c *gin.Context) (*presenter.UserCreateResponse, error)
	GetUser(c *gin.Context) (*presenter.UserGetResponse, error)
	UpdateUser(c *gin.Context) error
}

type userHandler struct {
	userUseCase   usecase.IUserUseCase
	userPresenter presenter.UserPresenter
	UserInput     input.CreateUserInput
}

func NewUserHandler(userUseCase usecase.IUserUseCase, userPresenter presenter.UserPresenter) IUserHandler {
	return &userHandler{userUseCase: userUseCase, userPresenter: userPresenter}
}

func (h *userHandler) CreateUser(ctx *gin.Context) (*presenter.UserCreateResponse, error) {
	req, err := h.userPresenter.CreateUserRequest(ctx)
	if err != nil || req == nil {
		return nil, err
	}
	input := input.CreateUserInput{Name: req.Name}
	encodeCtx := ctx.Request.Context()
	newCtx, err := h.userUseCase.CreateUser(encodeCtx, input)
	if err != nil {
		return nil, err
	}

	return h.userPresenter.CreateUserResponce(newCtx), nil

}

func (h *userHandler) GetUser(ctx *gin.Context) (*presenter.UserGetResponse, error) {
	logger := pkg.NewLogger()

	logger.Info("GetUser")
	req, err := h.userPresenter.GetUserRequest(ctx)
	if err != nil || req == nil {
		return nil, err
	}
	input := input.GetUserInput{UserID: req.UserID}
	newCtx, output, err := h.userUseCase.GetUser(ctx.Request.Context(), &input)
	if err != nil {
		return nil, err
	}

	return h.userPresenter.GetUserResponce(newCtx, output.Name), nil
}

func (h *userHandler) UpdateUser(ctx *gin.Context) error {
	logger := pkg.NewLogger()

	logger.Info("UpdateUser")
	req, err := h.userPresenter.UpdateUserRequest(ctx)
	if err != nil || req == nil {
		return err
	}
	input := input.UpdateUserInput{UserID: req.UserID, Name: req.Name}

	_, userCaseErr := h.userUseCase.UpdateUser(ctx.Request.Context(), &input)
	if userCaseErr != nil {
		return userCaseErr
	}

	return nil
}

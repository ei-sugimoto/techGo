package handler

import (
	"context"

	"github.com/ei-sugimoto/techGO/internal/handler/presenter"
	"github.com/ei-sugimoto/techGO/internal/usecase"
	input "github.com/ei-sugimoto/techGO/internal/usecase/Input"
	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	CreateUser(c *gin.Context) *presenter.UserCreateResponse
}

type userHandler struct {
	userUseCase   usecase.IUserUseCase
	userPresenter presenter.UserPresenter
	UserInput     input.UserInput
}

func NewUserHandler(userUseCase usecase.IUserUseCase, userPresenter presenter.UserPresenter) IUserHandler {
	return &userHandler{userUseCase: userUseCase, userPresenter: userPresenter}
}

func (h *userHandler) CreateUser(ctx *gin.Context) *presenter.UserCreateResponse {
	req, err := h.userPresenter.CreateUserRequest(ctx)
	if err != nil || req == nil {
		return h.userPresenter.CreateUserResponce(ctx, err)
	}
	input := input.UserInput{Name: req.Name}
	var encodeCtx context.Context
	encodeCtx = ctx.Request.Context()
	newCtx, err := h.userUseCase.CreateUser(encodeCtx, input)
	if err != nil {
		return h.userPresenter.CreateUserResponce(ctx, err)
	}

	return h.userPresenter.CreateUserResponce(newCtx, err)

}

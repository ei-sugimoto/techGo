package controller

import (
	"log/slog"
	"net/http"

	"github.com/ei-sugimoto/techGO/internal/adapter/presenter"
	"github.com/ei-sugimoto/techGO/internal/usecase"
	"github.com/ei-sugimoto/techGO/pkg"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase   usecase.UserUsecase
	UserPresenter *presenter.UserPresenter
	logger        *slog.Logger
}

func NewUserController(userUsecase usecase.UserUsecase, userPresenter *presenter.UserPresenter) *UserController {
	return &UserController{
		UserUsecase:   userUsecase,
		UserPresenter: userPresenter,
		logger:        pkg.NewLogger(),
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.logger.Error("failed to bind json")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	newCtx, err := c.UserUsecase.CreateUser(ctx, req.Name)
	if err != nil {
		c.logger.Error("failed to create user")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	res := c.UserPresenter.CreateUserResponce(newCtx)
	ctx.JSON(http.StatusOK, res)
	c.logger.Info("success to create user")
}

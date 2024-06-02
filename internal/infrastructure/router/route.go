package router

import (
	"github.com/ei-sugimoto/techGO/config"
	"github.com/ei-sugimoto/techGO/internal/adapter/controller"
	"github.com/ei-sugimoto/techGO/internal/adapter/presenter"
	"github.com/ei-sugimoto/techGO/internal/domain/service"
	"github.com/ei-sugimoto/techGO/internal/infrastructure/middleware"
	"github.com/ei-sugimoto/techGO/internal/infrastructure/repository"
	"github.com/ei-sugimoto/techGO/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowAllOrigins = true
	r.Use(cors.New(ginConfig))

	config.InitDB()

	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userUsecase := usecase.NewUserUsecase(userService)
	userPresenter := presenter.NewUserPresenter()
	userController := controller.NewUserController(userUsecase, userPresenter)

	r.POST("/user/create", middleware.Recovery(), middleware.NewUserAgent(), userController.CreateUser)

	return r
}

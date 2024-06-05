package router

import (
	"github.com/ei-sugimoto/techGO/internal/domain/service"
	"github.com/ei-sugimoto/techGO/internal/handler"
	"github.com/ei-sugimoto/techGO/internal/handler/middleware"
	"github.com/ei-sugimoto/techGO/internal/handler/presenter"
	"github.com/ei-sugimoto/techGO/internal/infrastructure/dao"
	"github.com/ei-sugimoto/techGO/internal/infrastructure/repository"
	"github.com/ei-sugimoto/techGO/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowAllOrigins = true
	ginConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	ginConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "X-Token"}
	r.Use(cors.New(ginConfig))

	db := dao.DataBase{}
	db.ConnectDataBase()

	userRepository := repository.NewUserRepository(&db)
	userService := service.NewUserService(userRepository)
	userUseCase := usecase.NewUserUsecase(userService)
	userPresenter := presenter.NewUserPresenter()
	userHandler := handler.NewUserHandler(userUseCase, *userPresenter)
	r.POST("/user/create", middleware.Recovery(), middleware.NewUserAgent(), func(c *gin.Context) {
		userHandlerCreate(userHandler, c)
	})
	r.GET("/user/get", middleware.Recovery(), middleware.NewUserAgent(), func(c *gin.Context) {
		userHandlerGet(userHandler, c)
	})

	return r
}

func userHandlerCreate(userHandler handler.IUserHandler, c *gin.Context) {
	res := userHandler.CreateUser(c)
	c.JSON(res.StatusCode, res)
}

func userHandlerGet(userHandler handler.IUserHandler, c *gin.Context) {
	res := userHandler.GetUser(c)
	c.JSON(res.StatusCode, res)
}

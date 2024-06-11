package router

import (
	"net/http"

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
	r.Use(ErrorHandler())
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

	userCharacterRepository := repository.NewUserCharacterRepository(&db)
	userCharacterService := service.NewUserCharacterService(userCharacterRepository)
	userCharacterUseCase := usecase.NewUserCharacterUseCase(userCharacterService)
	userCharacterPresenter := presenter.NewUserCharacterPresenter()
	userCharacterHandler := handler.NewUserCharacterHandler(userCharacterUseCase, *userCharacterPresenter)

	r.POST("/user/create", middleware.Recovery(), middleware.NewUserAgent(), func(c *gin.Context) {
		userHandlerCreate(userHandler, c)
	})
	r.GET("/user/get", middleware.Recovery(), middleware.NewUserAgent(), func(c *gin.Context) {
		userHandlerGet(userHandler, c)
	})
	r.PUT("/user/update", middleware.Recovery(), middleware.NewUserAgent(), func(c *gin.Context) {
		userHandlerUpdate(userHandler, c)
	})
	r.GET("/character/list", middleware.Recovery(), middleware.NewUserAgent(), func(c *gin.Context) {
		userCharacterHandlerGet(userCharacterHandler, c)
	})

	return r
}

func userHandlerCreate(userHandler handler.IUserHandler, c *gin.Context) {
	res, err := userHandler.CreateUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func userHandlerGet(userHandler handler.IUserHandler, c *gin.Context) {
	res, err := userHandler.GetUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"errors": c.Errors.Errors()})
		}
	}
}

func userHandlerUpdate(userHandler handler.IUserHandler, c *gin.Context) {
	err := userHandler.UpdateUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func userCharacterHandlerGet(userCharacterHandler handler.IUserCharacterHandler, c *gin.Context) {
	res, err := userCharacterHandler.GetUserCharacter(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

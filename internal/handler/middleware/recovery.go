package middleware

import (
	"log/slog"
	"net/http"

	"github.com/ei-sugimoto/techGO/pkg"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				pkg.NewLogger().Error("panic", slog.Any("error", err))
				c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			}
		}()
		c.Next()
	}
}

package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/ei-sugimoto/techGO/pkg"
	"github.com/gin-gonic/gin"
	"github.com/mileusna/useragent"
)

type key int

const userAgentKey key = iota

func GetUserAgent(r *http.Request) useragent.UserAgent {
	if r.Header.Get("User-Agent") == "" {
		return useragent.UserAgent{}
	}
	ua := useragent.Parse(r.Header.Get("User-Agent"))

	return ua
}

func NewUserAgent() gin.HandlerFunc {
	fn := func(r *http.Request) {
		ua := GetUserAgent(r)
		getLog := pkg.NewLogger().With(slog.String("path", "middleware/"))

		if ua.Desktop {
			getLog.Info("UserAgent", slog.String("OS", ua.OS), slog.String("Browser", ua.Name), slog.String("Version", ua.Version), slog.String("Device", "Desktop"))
		}
		if ua.Mobile {
			getLog.Info("UserAgent", slog.String("OS", ua.OS), slog.String("Browser", ua.Name), slog.String("Version", ua.Version), slog.String("Device", "Mobile"))
		}
		if ua.Tablet {
			getLog.Info("UserAgent", slog.String("OS", ua.OS), slog.String("Browser", ua.Name), slog.String("Version", ua.Version), slog.String("Device", "Tablet"))
		}
		r = r.WithContext(context.WithValue(r.Context(), userAgentKey, ua.OS))
	}
	return func(c *gin.Context) {
		fn(c.Request)
		c.Next()
	}
}

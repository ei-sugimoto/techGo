package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/ei-sugimoto/techGO/logger"
	"github.com/mileusna/useragent"
)

type key int
const userAgentKey key = iota

func  GetUserAgent(r *http.Request) useragent.UserAgent {
	if r.Header.Get("User-Agent") == "" {
		return useragent.UserAgent{}
	}
	ua := useragent.Parse(r.Header.Get("User-Agent"))

	return ua
}


func NewUserAgent(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ua := GetUserAgent(r)
		getLog := logger.NewLogger().With(slog.String("path", "middleware/"))
		getLog.Info("UserAgent OS:", slog.String("OS", ua.OS))
		getLog.Info("UserAgent Browser:", slog.String("Browser", ua.Name))
		getLog.Info("UserAgent Version:", slog.String("Version", ua.Version))
	
		
		if ua.Desktop {
			getLog.Info("UserAgent is Desktop")			
		}
		if ua.Mobile {
			getLog.Info("UserAgent is Mobile")
		}
		if ua.Tablet {
			getLog.Info("UserAgent is Tablet")
		}
		r = r.WithContext(context.WithValue(r.Context(), userAgentKey, ua.OS))
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
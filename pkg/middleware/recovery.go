package middleware

import (
	"log/slog"
	"net/http"

	"github.com/ei-sugimoto/techGO/internal/pkg"
)

func Recovery(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		newLogger := pkg.NewLogger().With(slog.String("path", "middleware/"))
		defer func() {
			if r := recover(); r != nil {
				newLogger.Error("panic", slog.Any("panic", r))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			}
		}()

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

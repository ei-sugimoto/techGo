package middleware

import (
	"fmt"
	"net/http"
)

func Recovery(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// TODO: ここに実装をする
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("get panic %v\n", r)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			}
		}()

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

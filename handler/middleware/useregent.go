package middleware

import (
	"context"
	"log"
	"net/http"

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
		log.Println("UserAgent OS:", ua.OS)
		log.Println("UserAgent Browser:", ua.Name)
		log.Println("UserAgent Version:", ua.Version)
		if ua.Desktop {
			log.Println("UserAgent is Desktop")
		}
		if ua.Mobile {
			log.Println("UserAgent is Mobile")
		}
		if ua.Tablet {
			log.Println("UserAgent is Tablet")
		}
		r = r.WithContext(context.WithValue(r.Context(), userAgentKey, ua.OS))
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
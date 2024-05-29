package router

import (
	"database/sql"
	"net/http"

	"github.com/ei-sugimoto/techGO/handler"
	"github.com/ei-sugimoto/techGO/handler/middleware"
	"github.com/ei-sugimoto/techGO/service"
	"github.com/rs/cors"
)

func NewRouter(db *sql.DB) http.Handler {
	mux := http.NewServeMux()
	userService := service.NewUser(db)
	userHandler := handler.NewUserHandler(userService)
	mux.Handle("/user/create", http.HandlerFunc(userHandler.CreateUser))
	mux.Handle("/user/get", http.HandlerFunc(userHandler.GetUser))
	mux.Handle("/user/update", http.HandlerFunc(userHandler.UpdateUser))
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("panic test")
	})

	RecoveryMux := middleware.Recovery(mux)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})

	userAgentMux := middleware.NewUserAgent(RecoveryMux)
	handler := c.Handler(userAgentMux)

	return handler
}

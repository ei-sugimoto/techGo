package router

import (
	"database/sql"
	"net/http"

	"github.com/ei-sugimoto/techGO/handler"
	"github.com/ei-sugimoto/techGO/handler/middleware"
	"github.com/ei-sugimoto/techGO/service"
)


func NewRouter(db *sql.DB) http.Handler {
	mux := http.NewServeMux()
	userService := service.NewUserService(db)
	userHandler := handler.NewUserHandler(userService)
	mux.Handle("/users", userHandler)
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("panic test")
	})
	RecoveryMux := middleware.Recovery(mux)
	return RecoveryMux
}
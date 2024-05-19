package router

import (
	"database/sql"
	"net/http"

	"github.com/ei-sugimoto/techGO/handler/middleware"
)


func NewRouter(db *sql.DB) http.Handler {
	mux := http.NewServeMux()
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
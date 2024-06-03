package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ei-sugimoto/techGO/internal/handler/router"
	"github.com/ei-sugimoto/techGO/pkg"
)

func main() {

	r := router.NewRouter()
	logger := pkg.NewLogger()
	go func() {
		if err := r.Run(":8080"); err != nil {
			logger.Error("Server start error")
			return
		}
	}()
	logger.Info("Server started on :8080")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	logger.Info("Shutting down server...")
}

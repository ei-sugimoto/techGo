package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ei-sugimoto/techGO/connect"
	"github.com/ei-sugimoto/techGO/handler/router"
	_ "github.com/go-sql-driver/mysql"
)


func main()  {
	const (
		dbDriver = "mysql"
	)
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName
	db, err := connect.SetupDB(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
		fmt.Println("DB接続エラー")
		return
	}
	defer db.Close()
	fmt.Println("DB接続成功")
	mux := router.NewRouter(db)
	srv := &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }
	go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()
	log.Println("Server started on :8080")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Println("Shutting down server...")

	
}
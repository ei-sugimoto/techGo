package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	http.ListenAndServe(":8080", mux)
}
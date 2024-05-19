package main

import (
	"time"
)


func setupDB() error {
	const (
		host	 = "localhost"
		port    = 8080
	)

	var err error
	time.Local, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}
	DB, err = connectdb.setupDB()
	return nil
}
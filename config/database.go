package config

import (
	"github.com/ei-sugimoto/techGO/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	const (
		dbDriver = "mysql"
	)
	host, port, user, password, dbName := Env()

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	logger := pkg.NewLogger()
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Error(("Failed to connect to database"))
	}
	logger.Info("Connected to database")

	return
}

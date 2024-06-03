package dao

import (
	"github.com/ei-sugimoto/techGO/config"
	"github.com/ei-sugimoto/techGO/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DataBase struct {
	GormDB *gorm.DB
}

func NewDataBase(db *gorm.DB) *DataBase {
	return &DataBase{GormDB: db}
}

func (d *DataBase) ConnectDataBase() {
	const (
		dbDriver = "mysql"
	)
	host, port, user, password, dbName := config.Env()

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	logger := pkg.NewLogger()

	var err error
	d.GormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to database")
		return
	}
	logger.Info("Connected to database")

	return
}

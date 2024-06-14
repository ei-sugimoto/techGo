package dao

import (
	"fmt"

	"github.com/ei-sugimoto/techGO/config"
	"github.com/ei-sugimoto/techGO/internal/domain/model"
	"github.com/ei-sugimoto/techGO/pkg"
	"github.com/google/uuid"
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

	host, port, user, password, dbName := config.Env()

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	logger := pkg.NewLogger()

	var err error
	d.GormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to database")
		panic(err)
	}
	d.Drop()
	d.Migrate()
	d.Seed()
	logger.Info("Connected to database")

}

func (d *DataBase) Migrate() {
	d.GormDB.AutoMigrate(&model.User{})
	d.GormDB.AutoMigrate(&model.Character{})
	d.GormDB.AutoMigrate(&model.UserCharacter{})
}

func (d *DataBase) Drop() {
	d.GormDB.Migrator().DropTable(&model.UserCharacter{}, &model.Character{}, &model.User{})
}

func (d *DataBase) Seed() {
	var characters []*model.Character
	for i := 1; i <= 15; i++ {
		character, err := model.NewCharacter(uuid.New().String(), fmt.Sprintf("Character%d", i), i)
		if err != nil {
			panic(err)
		}
		characters = append(characters, character)
	}
	raretestCharacter, err := model.NewCharacter(uuid.New().String(), "RaretestCharacter", 100)
	if err != nil {
		panic(err)
	}
	characters = append(characters, raretestCharacter)
	d.GormDB.Create(&characters)

}

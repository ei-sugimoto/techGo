package dao

import (
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
	userId1 := uuid.New().String()
	userId2 := uuid.New().String()
	users := []model.User{
		{
			UserID: userId1,
			Name:   "Alice",
		},
		{
			UserID: userId2,
			Name:   "Bob",
		},
	}
	d.GormDB.Create(&users)
	characterId1 := uuid.New().String()
	characterId2 := uuid.New().String()
	characters := []model.Character{
		{
			CharacterID: characterId1,
			Name:        "Warrior",
		},
		{
			CharacterID: characterId2,
			Name:        "Magician",
		},
		{
			CharacterID: uuid.New().String(),
			Name:        "Archer",
		},
		{
			CharacterID: uuid.New().String(),
			Name:        "Priest",
		},
	}
	d.GormDB.Create(&characters)
	userCharacterId1 := uuid.New().String()
	userCharacterId2 := uuid.New().String()
	userCharacters := []model.UserCharacter{
		{
			UserCharacterID: userCharacterId1,
			CharacterID:     characterId1,
			UserID:          userId1,
		},
		{
			UserCharacterID: userCharacterId2,
			CharacterID:     characterId2,
			UserID:          userId2,
		},
	}
	d.GormDB.Create(&userCharacters)

}

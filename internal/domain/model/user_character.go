package model

import (
	"github.com/google/uuid"
)

type UserCharacter struct {
	UserCharacterID uuid.UUID `gorm:"type:uuid;primary_key;"`
	Character       Character `gorm:"foreignkey:CharacterID;association_foreignkey:CharacterID;association_autoupdate:false;association_autocreate:false;"`
	User            User      `gorm:"foreignkey:UserID;association_foreignkey:UserID;association_autoupdate:false;association_autocreate:false;"`
	Name            string    `gorm:"type:varchar(255);not null;"`
}

func (uc *UserCharacter) TableName() string {
	return "user_character"
}

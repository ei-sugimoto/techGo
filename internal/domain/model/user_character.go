package model

import (
	"github.com/google/uuid"
)

type UserCharacter struct {
	UserCharacterID uuid.UUID `gorm:"type:uuid;primary_key;column:user_character_id"`
	CharacterID     uuid.UUID `gorm:"type:uuid;not null"`
	UserID          uuid.UUID `gorm:"type:uuid;not null"`
	User            User      `gorm:"foreignKey:UserID;association_foreignkey:UserID"`
	Character       Character `gorm:"foreignKey:CharacterID;association_foreignkey:CharacterID"`
}

func (uc *UserCharacter) TableName() string {
	return "user_character"
}

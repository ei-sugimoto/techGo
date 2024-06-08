package model

import (
	"github.com/google/uuid"
)

type UserCharacter struct {
	UserCharacterID uuid.UUID `gorm:"type:uuid;primary_key;"`
	User            User      `gorm:"foreignkey:UserID;association_foreignkey:UserID;association_autoupdate:false;association_autocreate:false;"`
}

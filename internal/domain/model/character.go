package model

import "github.com/google/uuid"

type Character struct {
	CharacterID uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name        string    `gorm:"type:varchar(255);not null;"`
}

func (c *Character) TableName() string {
	return "character"
}

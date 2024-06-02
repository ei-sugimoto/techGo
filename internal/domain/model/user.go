package model

import (
	"github.com/google/uuid"
)

type User struct {
	UserID uuid.UUID `gorm:"type:uuid;primary_key;unique;not null;"`
	Name   string    `gorm:"type:varchar(255);not null;"`
}

func (u *User) TableName() string {
	return "user"
}

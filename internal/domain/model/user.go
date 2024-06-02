package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID   uuid.UUID `gorm:"type:uuid;primary_key;unique;not null;"`
	Name     string    `gorm:"type:varchar(255);not null;"`
	CreateAt time.Time `gorm:"type:timestamp;not null;"`
	UpdateAt time.Time `gorm:"type:timestamp;not null;"`
}

func (u *User) BeforeCreate() {
	u.CreateAt = time.Now()
	u.UpdateAt = time.Now()
}

func (u *User) BeforeUpdate() {
	u.UpdateAt = time.Now()
}

func (u *User) TableName() string {
	return "users"
}

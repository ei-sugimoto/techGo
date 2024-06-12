package model

type User struct {
	UserID string `gorm:"type:char(36);primary_key;unique;not null;"`
	Name   string `gorm:"type:varchar(255);not null;"`
}

func (u *User) TableName() string {
	return "user"
}

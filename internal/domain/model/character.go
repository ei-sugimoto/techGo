package model

type Character struct {
	CharacterID string `gorm:"type:char(36);primary_key;unique;not null;"`
	Name        string `gorm:"type:varchar(255);not null;"`
}

func (c *Character) TableName() string {
	return "character"
}

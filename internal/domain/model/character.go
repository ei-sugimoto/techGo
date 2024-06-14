package model

type Character struct {
	CharacterID string  `gorm:"type:char(36);primary_key;unique;not null;"`
	Name        string  `gorm:"type:varchar(255);not null;"`
	Probability float64 `gorm:"default:50;"`
}

func (c *Character) TableName() string {
	return "character"
}

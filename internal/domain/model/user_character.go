package model

type UserCharacter struct {
	UserCharacterID string `gorm:"type:char(36);primary_key;not null;"`
	CharacterID     string
	UserID          string
	User            User      `gorm:"references:UserID"`
	Character       Character `gorm:"references:CharacterID"`
}

func (uc *UserCharacter) TableName() string {
	return "user_character"
}

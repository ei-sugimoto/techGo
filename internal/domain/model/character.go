package model

import "errors"

type Character struct {
	CharacterID string `gorm:"type:char(36);primary_key;unique;not null;"`
	Name        string `gorm:"type:varchar(255);not null;"`
	Rarity      int    `gorm:"type:int;not null;"`
}

func (c *Character) TableName() string {
	return "character"
}

func NewCharacter(characterID, name string, rarity int) (*Character, error) {
	if err := isValidRarity(rarity); err != nil {
		return nil, err
	}
	return &Character{
		CharacterID: characterID,
		Name:        name,
		Rarity:      rarity,
	}, nil
}

func isValidRarity(rarity int) error {
	if rarity >= 1 && rarity <= 5 {
		return nil
	} else {
		return errors.New("invalid rarity")
	}
}

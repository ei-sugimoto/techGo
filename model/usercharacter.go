package model

type (
	UserCharacter struct {
		UserCharacterID string `json:"user_character_id"`
		CharacterID string `json:"character_id"`
		Name string `json:"name"`
	}

	GetUserCharacterRequest struct{
	}
	GetUserCharacterResponce struct{
		UserCharacters []*UserCharacter `json:"UserCharacter"`
	}
)


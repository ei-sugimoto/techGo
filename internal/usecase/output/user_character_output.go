package output

type GetUserCharacterOutput struct {
	UserCharacterID string
	CharacterID     string
	Name            string
}

type GetUserCharacterOutputs []GetUserCharacterOutput

type CreateUserCharacterOutput struct {
	UserCharacterID string
	CharacterID     string
	Name            string
}

type CreateUserCharacterOutputs []CreateUserCharacterOutput

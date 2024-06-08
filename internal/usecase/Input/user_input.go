package input

type CreateUserInput struct {
	Name string
}

type GetUserInput struct {
	UserID string
}

type UpdateUserInput struct {
	UserID string
	Name   string
}

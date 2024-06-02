package presenter

import "context"

type UserPresenter struct{}

func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}

type UserCreateResponse struct {
	Token string `json:"token"`
}

func (p *UserPresenter) CreateUserResponce(ctx context.Context) *UserCreateResponse {
	token := ctx.Value("token").(string)
	return &UserCreateResponse{
		Token: token,
	}
}

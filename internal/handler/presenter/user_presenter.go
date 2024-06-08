package presenter

import (
	"context"

	"github.com/ei-sugimoto/techGO/pkg"
	"github.com/gin-gonic/gin"
)

type UserPresenter struct{}

func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}

type UserCreateResponse struct {
	Token string `json:"token"`
}

type UserCreateRequest struct {
	Name string `json:"name"`
}

type UserGetRequest struct {
	UserID string `json:"user_id"`
}

type UserGetResponse struct {
	Name string `json:"name"`
}

type UserUpdateRequest struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}

type UserUpdateResponse struct {
}

func (p *UserPresenter) CreateUserResponce(ctx context.Context) *UserCreateResponse {
	token := ctx.Value("token").(string)
	return &UserCreateResponse{
		Token: token,
	}
}

func (p *UserPresenter) CreateUserRequest(ctx *gin.Context) (*UserCreateRequest, error) {
	var req UserCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (p *UserPresenter) GetUserRequest(ctx *gin.Context) (*UserGetRequest, error) {
	var req UserGetRequest
	token := ctx.GetHeader("x-token")
	userId, err := pkg.DecodeJwt(token, "user_id")
	if err != nil {
		return nil, err
	}
	req.UserID = userId

	return &req, nil
}

func (p *UserPresenter) GetUserResponce(ctx context.Context, name string) *UserGetResponse {

	return &UserGetResponse{
		Name: name,
	}
}

func (p *UserPresenter) UpdateUserRequest(ctx *gin.Context) (*UserUpdateRequest, error) {
	type Request struct {
		Name string `json:"name"`
	}
	var body Request
	if err := ctx.BindJSON(&body); err != nil {
		// エラーハンドリング
		return nil, err
	}

	var req UserUpdateRequest
	req.Name = body.Name
	token := ctx.GetHeader("x-token")
	userId, err := pkg.DecodeJwt(token, "user_id")
	if err != nil {
		return nil, err
	}
	req.UserID = userId

	return &req, nil
}

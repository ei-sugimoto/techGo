package presenter

import (
	"context"
	"net/http"

	"github.com/ei-sugimoto/techGO/pkg"
	"github.com/gin-gonic/gin"
)

type UserPresenter struct{}

func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}

type UserCreateResponse struct {
	Token      string `json:"token"`
	StatusCode int    `json:"code"`
}

type UserCreateRequest struct {
	Name string `json:"name"`
}

type UserGetRequest struct {
	UserID string `json:"user_id"`
}

type UserGetResponse struct {
	Name       string `json:"name"`
	StatusCode int    `json:"code"`
}

func (p *UserPresenter) CreateUserResponce(ctx context.Context, err error) *UserCreateResponse {
	if err != nil {
		return &UserCreateResponse{
			Token:      "",
			StatusCode: http.StatusBadRequest,
		}
	}
	token := ctx.Value("token").(string)
	return &UserCreateResponse{
		Token:      token,
		StatusCode: http.StatusOK,
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

func (p *UserPresenter) GetUserResponce(ctx context.Context, err error, name string) *UserGetResponse {
	if err != nil {
		return &UserGetResponse{
			Name:       "",
			StatusCode: http.StatusBadRequest,
		}
	}
	return &UserGetResponse{
		Name:       name,
		StatusCode: http.StatusOK,
	}
}

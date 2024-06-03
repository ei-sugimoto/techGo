package presenter

import (
	"context"
	"net/http"

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

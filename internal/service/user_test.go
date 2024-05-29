package service_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/ei-sugimoto/techGO/model"
	"github.com/ei-sugimoto/techGO/service"
	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func TestCreateUser_EmptyName(t *testing.T) {
	res, err := service.NewUser(nil).CreateUser(context.TODO(), &model.UserCreateRequest{Name: ""})
	if err == nil {
		t.Errorf("expected error, but got nil")
	} else {
		if err.Message != "name is required" {
			t.Errorf("expected error message is 'name is required', but got %v", err.Error())
		}

		if err.Code != http.StatusBadRequest {
			t.Errorf("expected error code is %v, but got %v", http.StatusBadRequest, err.Code)
		}
	}
	if res != nil {
		t.Errorf("expected nil, but got %v", res)
	}
}

func TestGetUser_Invalid_Token(t *testing.T) {
	token := "invalid_token"
	res, err := service.NewUser(nil).GetUser(context.TODO(), token)

	if err == nil {
		t.Errorf("expected error, but got nil")
	}
	if res != nil {
		t.Errorf("expected nil, but got %v", res)
	}

}

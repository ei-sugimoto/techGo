package service_test

import (
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
	res, err := service.NewUserCharacter(nil).CreateUser(nil, &model.UserCreateRequest{Name: ""})
	if err == nil {
		t.Errorf("expected error, but got nil")
	}
	if res != nil {
		t.Errorf("expected nil, but got %v", res)
	}

	if err.Message != "name is required" {
		t.Errorf("expected error message is 'name is required', but got %v", err.Error())
	}

	if err.Code != http.StatusBadRequest{
		t.Errorf("expected error code is %v, but got %v", http.StatusBadRequest, err.Code)
	}
}

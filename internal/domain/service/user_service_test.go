package service_test

import (
	"context"
	"log"
	"testing"

	"github.com/ei-sugimoto/techGO/internal/domain/model"
	"github.com/ei-sugimoto/techGO/internal/domain/repository/mock_repository"
	"github.com/ei-sugimoto/techGO/internal/domain/service"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockIUserRepository(ctrl)
	mockUUID := uuid.New()
	mockUser := &model.User{UserID: mockUUID.String(), Name: "testName"}
	mockRepo.EXPECT().GetUser(gomock.Any(), "testID").Return(mockUser, nil)

	userService := service.NewUserService(mockRepo)

	_, user, err := userService.GetUser(context.Background(), "testID")
	log.Println(mockUser.UserID, user.UserID)
	assert.Equal(t, mockUser, user)
	assert.Nil(t, err)

}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockIUserRepository(ctrl)
	mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)

	userService := service.NewUserService(mockRepo)

	_, err := userService.CreateUser(context.Background(), "testName")

	assert.Nil(t, err)
}

package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ei-sugimoto/techGO/model"
	"github.com/ei-sugimoto/techGO/service"
)

type (
	UserHandler struct {
		userService *service.User
	}
)

func NewUserHandler(userService *service.User) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUsers(ctx context.Context, req *model.GetUserRequest) (*model.GetUserResponce, error) {
	res, err := h.userService.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return &model.GetUserResponce{Users: res}, nil

}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	req := &model.UserCreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Name == "" {
		log.Printf("name is Required code:%d\n", http.StatusBadRequest)
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}
	res, err := h.userService.CreateUser(r.Context(), req)
	if err != nil {
		http.Error(w, err.Message, err.Code)
		return
	}
	jsonRes, _ := json.Marshal(res.Token)
	if jsonRes == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(jsonRes)
	log.Println("User created")
}

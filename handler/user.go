package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ei-sugimoto/techGO/model"
	"github.com/ei-sugimoto/techGO/service"
)

type (
	UserHandler struct{
		userService *service.UserService
	}
)

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUsers(ctx context.Context, req *model.GetUsersRequest)(*model.GetUsersResponce, error) {
	users, err := h.userService.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return &model.GetUsersResponce{Users: users}, nil

}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		req := &model.GetUsersRequest{}
		res, err := h.GetUsers(r.Context(), req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonRes , err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonRes)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
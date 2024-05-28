package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ei-sugimoto/techGO/logger"
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

func (h *UserHandler) GetLogger() *slog.Logger {
	return logger.NewLogger().With(slog.String("path", "handler/"))
}


func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.GetLogger().Error("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	req := &model.UserCreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Name == "" {
		h.GetLogger().Error("name is required", slog.Int("code", http.StatusBadRequest))
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
	h.GetLogger().Info("User created")
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodGet {
		h.GetLogger().Error("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get("X-Token")

	if token == "" {
		h.GetLogger().Error("Token is required")
		http.Error(w, "Token is required", http.StatusBadRequest)
		return
	}

	res, err := h.userService.GetUser(r.Context(), token)
	if err != nil {
		http.Error(w, err.Message, err.Code)
		return
	}
	jsonRes, _ := json.Marshal(res)
	if jsonRes == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(jsonRes)
	h.GetLogger().Info("User Get")

}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
	h.GetLogger().Error("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	req := &model.UserUpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Name == "" {
		h.GetLogger().Error("name is required", slog.Int("code", http.StatusBadRequest))
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	token := r.Header.Get("X-Token")

	if token == "" {
		h.GetLogger().Error("Token is required")
		http.Error(w, "Token is required", http.StatusBadRequest)
		return
	}

	res, err := h.userService.UpdateUser(r.Context(), req, token)
	if err != nil {
		http.Error(w, err.Message, err.Code)
		return
	}
	jsonRes, _ := json.Marshal(res)
	if jsonRes == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(jsonRes)
	h.GetLogger().Info("User updated")
}

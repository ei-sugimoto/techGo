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
	UserHandler struct{
		userService *service.UserCharacter
	}
)

func NewUserHandler(userService *service.UserCharacter) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUserCharacters(ctx context.Context, req *model.GetUserCharacterRequest)(*model.GetUserCharacterResponce, error) {
	res, err := h.userService.GetUserCharacters(ctx)
	if err != nil {
		return nil, err
	}

	return &model.GetUserCharacterResponce{UserCharacters: res}, nil

}

func (h *UserHandler) CreateUser(ctx context.Context, req *model.UserCreateRequest)(*model.UserCreateResponse, *model.CutomError) {
	res, err := h.userService.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return &model.UserCreateResponse{Token: res.Token}, nil
}


func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:{
		req := &model.UserCreateRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.Name == "" {
			http.Error(w, "name is required", http.StatusBadRequest)
			return
		}
		res, err := h.CreateUser(r.Context(), req)
		if err != nil {
			http.Error(w, err.Message, err.Code)
			return
		}
		jsonRes , _ := json.Marshal(res.Token)
		if(jsonRes == nil){
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write(jsonRes)
		log.Println("User created")
	}

		
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
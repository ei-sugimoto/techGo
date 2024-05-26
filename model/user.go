package model

type (
	User struct {
		UserID string `json:"user_id"`
		Name   string `json:"name"`
	}

	GetUserRequest struct {
	}
	GetUserResponce struct {
		Users []*User `json:"User"`
	}

	UserCreateRequest struct {
		Name string `json:"name"`
	}
	UserCreateResponse struct {
		Token string `json:"token"`
	}

	UserGetResponse struct {
		Name string `json:"name"`
	}

	UserUpdateRequest struct {
		Name string `json:"name"`
	}

	UserUpdateResponse struct {
		Token string `json:"token"`
	}
)

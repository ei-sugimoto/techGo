package model

type (
	User struct {
		ID 	 int    `json:"id"`
		UserName string `json:"username"`
		Password string `json:"password"`
		Email	string `json:"email"`
		CreatedAt string `json:"created_at"`
	}

	GetUsersRequest struct{
	}
	GetUsersResponce struct{
		Users []*User `json:"users"`
	}
)


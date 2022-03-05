package handler

import "github.com/auto-crud/models"

type userResponse struct {
	User struct {
		Id         uint   `json:"id"`
		Name       string `json:"name"`
		SecondName string `json:"second_name"`
		Online     bool   `json:"online"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at"`
	} `json:"user"`
}

func newUserResponse(u *models.SelectUser) *userResponse {
	r := new(userResponse)
	r.User.Id = u.Id
	r.User.Name = u.Name
	r.User.SecondName = u.SecondName
	r.User.Online = u.Online
	r.User.CreatedAt = u.CreatedAt
	r.User.UpdatedAt = u.UpdatedAt
	return r
}

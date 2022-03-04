package models

type User struct {
	Name       string `json:"name"`
	SecondName string `json:"second_name"`
	Online     bool   `json:"online"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

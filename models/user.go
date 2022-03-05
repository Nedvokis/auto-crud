package models

type InsertUser struct {
	Name       string `json:"name"`
	SecondName string `json:"second_name"`
	Online     bool   `json:"online"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
type SelectUser struct {
	Id         uint   `json:"id" query:"id"`
	Name       string `json:"name" query:"name"`
	SecondName string `json:"second_name" query:"second_name"`
	Online     bool   `json:"online" query:"online"`
	CreatedAt  string `json:"created_at" query:"created_at"`
	UpdatedAt  string `json:"updated_at" query:"updated_at"`
}

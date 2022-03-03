package models

import "time"

type User struct {
	Id         uint      `json:"id"`
	Name       string    `json:"name"`
	SecondName string    `json:"second_name"`
	Online     bool      `json:"online"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

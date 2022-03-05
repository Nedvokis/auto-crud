package handler

import (
	"github.com/labstack/echo/v4"
)

type GetUserRequest struct {
	User struct {
		Id        uint   `json:"id" query:"id" validate:"required"`
		TableName string `json:"table_name" query:"table_name" validate:"required"`
	} `json:"user"`
}

func (r *GetUserRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

package handler

import (
	"net/http"

	"github.com/auto-crud/models"
	"github.com/auto-crud/utils"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUser(c echo.Context) error {
	var u models.SelectUser
	req := &GetUserRequest{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	h.store.GetOne(req.User.TableName, req.User.Id, &u)

	return c.JSON(http.StatusCreated, newUserResponse(&u))
}

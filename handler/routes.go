package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	guestUsers := v1.Group("user")
	guestUsers.GET("", h.GetUser)
	// guestUsers.GET("/login")
}

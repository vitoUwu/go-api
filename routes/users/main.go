package users

import (
	"github.com/labstack/echo/v4"
)

type UserBind struct {
	Username string `json:"username" validate:"required"`
}

func RegisterRoutes(e *echo.Echo) {
	router := e.Group("/users")
	router.GET("/:username", findUserByUsername)
	router.POST("/", createUser)
	router.PATCH("/:username", updateUsername)
}

package interfaces

import (
	"github.com/labstack/echo/v4"
)

type Controllers struct {
	userController *UserController
}

func NewControllers(
	userController *UserController,
) *Controllers {
	return &Controllers{
		userController: userController,
	}
}

func (c *Controllers) Mount(group *echo.Group) {
	c.userController.Mount(group.Group("/user"))
}

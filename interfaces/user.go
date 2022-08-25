package interfaces

import (
	"go_api_sample/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(
	userUsecase usecase.UserUsecase,
) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (c *UserController) Mount(group *echo.Group) {
	group.GET("/:id", c.Show)
	group.POST("/", c.Create)
	group.PUT("/", c.Update)
	group.DELETE("/:id", c.Delete)
}

func (c *UserController) Show(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user, err := c.userUsecase.FindUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return e.JSON(http.StatusOK, user)
}

func (c *UserController) Create(e echo.Context) error {
	request := &struct {
		Name string
	}{}
	if err := e.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err := c.userUsecase.CreateUser(
		&usecase.CreateUserInput{
			Name: request.Name,
		},
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return e.String(http.StatusOK, "status ok")
}

func (c *UserController) Update(e echo.Context) error {
	request := &struct {
		ID   int
		Name string
	}{}
	if err := e.Bind(request); err != nil {
		return err
	}

	err := c.userUsecase.UpdateUser(
		&usecase.UpdateUserInput{
			ID:   request.ID,
			Name: request.Name,
		},
	)
	if err != nil {
		return err
	}

	return e.String(http.StatusOK, "status ok")
}

func (c *UserController) Delete(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := c.userUsecase.DeleteUser(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return e.String(http.StatusOK, "status ok")
}

package controllers

import (
	"github.com/goki0524/clean-arch/usecases"
	"github.com/goki0524/clean-arch/usecases/dto/usersdto"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type UsersController struct {
	usersUseCase usecases.IUsersUseCase
}

func NewUsersController(
	usersUseCase usecases.IUsersUseCase,
) *UsersController {
	return &UsersController{
		usersUseCase: usersUseCase,
	}
}

func (uc *UsersController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.QueryParam("id")

		if id == "" {
			return c.JSON(http.StatusBadRequest, "パラメータを正しく設定してください")
		}

		idUint, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			logrus.Errorf("Error strconv: %v", err)
		}

		in := usersdto.GetInput{
			ID: uint(idUint),
		}
		out := uc.usersUseCase.Get(in)

		return c.JSON(http.StatusOK, out)
	}
}

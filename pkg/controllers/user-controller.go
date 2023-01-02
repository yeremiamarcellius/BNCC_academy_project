package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/youhane/bncc_academy_final/pkg/models"
)

type UserController struct{}

func (uc *UserController) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := models.GetUserByEmail(email)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	if user.Password != password {
		return c.JSON(http.StatusUnauthorized, "Wrong password")
	}

	token, err := models.GenerateJWT(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token":    token,
		"username": user.Username,
		"email":    user.Email,
	})
}

func (uc *UserController) Register(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	user := models.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	user.CreateUser()

	return c.JSON(http.StatusOK, user)
}

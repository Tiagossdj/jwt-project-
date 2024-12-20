package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Login é o handler para o endPoint /auth/login
func Login(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Login realizado com sucesso!",
	})
}

// Register é o handler para o endPoint /auth/register
func Register(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]string{
		"Message": "Registro realizado com sucesso!",
	})

}

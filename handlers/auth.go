package handlers

import (
	"net/http"
	"time"

	"github.com/Tiagossdj/jwt-project-/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
)

// chave secreta para assinar o token jwt
var jwtSecret = []byte("your secret key!")

// Login é o handler para o endPoint /auth/login
func Login(c echo.Context) error {
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Message{
			Message: "Invalid data!",
		})
	}

	// validação básica para substituir com a lógica real de autenticação depois
	if req.Email != "---" || req.Password != "---" {
		return c.JSON(http.StatusUnauthorized, model.Message{
			Message: "invalid credentials!",
		})
	}

	// geração de token JWT

	claims := &jwt.RegisteredClaims{
		Subject:   req.Email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 24h de duração.
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Message{
			Message: "Error to generate Token!",
		})
	}
	return c.JSON(http.StatusOK, model.Message{
		Message: "Token:" + signedToken,
	})
}

// Register é o handler para o endPoint /auth/register
func Register(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]string{
		"Message": "Registro realizado com sucesso!",
	})

}

package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/Tiagossdj/jwt-project-/handlers"
	"github.com/Tiagossdj/jwt-project-/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Recuperar o header Authorization
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, model.Message{
				Message: "Missing or malformed jwt",
			})
		}

		// Verifica se o token está no formato Bearer <token>
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, model.Message{
				Message: "Missing or malformed jwt",
			})
		}

		// Remove o prefixo "Bearer" e pega o token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// log.Printf("Validating Token: %s", tokenString) // Log para depuração!

		// Parse do token com claims
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			// Validar o método de assinatura e retornar a chave secreta
			return []byte(handlers.JwtSecret), nil
		})

		if err != nil {
			log.Printf("Error parsing token: %v", err) // Log de erro
			return c.JSON(http.StatusUnauthorized, model.Message{
				Message: "Invalid Token",
			})
		}

		// Verifica se o token é válido
		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, model.Message{
				Message: "Invalid Token",
			})
		}

		// Adiciona o usuário no contexto
		c.Set("user", claims)

		return next(c)
	}
}

package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/v4/handlers"
)

func InitRoutes(e *echo.Echo) {
	//Rota de teste
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	//Rotas para Autenticação
	authGroup := e.Group("/auth")
	authGroup.POST("/login", handlers.Login)
	authGroup.POST("/register", handlers.Register)

}

// Grupo de rotas protegidas (exemplo para proteger com JWT futuramente)
// protectedGroup := e.Group("/protected")
// protectedGroup.Use(middleware.JWT([]byte("secret")))
// protectedGroup.GET("/dashboard", handlers.Dashboard)

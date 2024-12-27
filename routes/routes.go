package routes

import (
	"github.com/Tiagossdj/jwt-project-/db"
	"github.com/Tiagossdj/jwt-project-/handlers"
	"github.com/labstack/echo"
)

func InitRoutes(e *echo.Echo) {
	//Rota de teste
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	//Rotas para Autenticação
	authGroup := e.Group("/auth")

	authGroup.POST("/login", handlers.Login)

	dbConn, _ := db.ConnDB()
	authGroup.POST("/auth/register", func(c echo.Context) error {
		return handlers.Register(c, dbConn)
	})

}

// Grupo de rotas protegidas (exemplo para proteger com JWT futuramente)
// protectedGroup := e.Group("/protected")
// protectedGroup.Use(middleware.JWT([]byte("secret")))
// protectedGroup.GET("/dashboard", handlers.Dashboard)

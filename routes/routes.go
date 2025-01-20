package routes

import (
	"log"

	"github.com/Tiagossdj/jwt-project-/db"
	"github.com/Tiagossdj/jwt-project-/handlers"
	"github.com/Tiagossdj/jwt-project-/middlewares"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	//Rota de teste
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	//Rotas para Autenticação
	authGroup := e.Group("/auth")

	// Grupo de rotas protegidas (para proteger com JWT)
	protectedGroup := e.Group("/protected")
	protectedGroup.Use(middlewares.JwtMiddleware)

	dbConn, err := db.ConnDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	authGroup.POST("/login", func(c echo.Context) error {
		return handlers.Login(c, dbConn)
	})

	authGroup.POST("/register", func(c echo.Context) error {
		return handlers.Register(c, dbConn)
	})

	protectedGroup.GET("/profile", func(c echo.Context) error {
		return handlers.GetProfile(c)
	})

}

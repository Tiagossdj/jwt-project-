package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/Tiagossdj/jwt-project-/db"
	_ "github.com/Tiagossdj/jwt-project-/docs"
	"github.com/Tiagossdj/jwt-project-/routes"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title JWT authentication API
// @version 1.0
// @description This API handles user authentication using JWT with Echo framework.
// @host localhost:8888
// @BasePath /
// @schemes http
func main() {

	e := echo.New()

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8888"))

	//DB
	conn, err := db.ConnDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close()

	//Rotas
	routes.InitRoutes(e)

	//Server
	if err := e.Start(":8888"); err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Error to start Server:%v", err)
	}

}

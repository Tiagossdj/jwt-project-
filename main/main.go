package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/Tiagossdj/jwt-project-/db"
	"github.com/Tiagossdj/jwt-project-/routes"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

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

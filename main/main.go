package main

import (
	"errors"
	//"jwt-project/routes"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/v4/routes"
)

func main() {
	e := echo.New()

	//Routes

	//login
	//e.POST("/")

	//register
	//e.POST("/")

	//profile
	//e.GET("/")

	routes.InitRoutes(e)

	//Server
	if err := e.Start(":8080"); err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Error to start Server:%v", err)
	}

}

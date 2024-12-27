package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/Tiagossdj/jwt-project-/model"
	"github.com/Tiagossdj/jwt-project-/routes"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	//Routes

	//login
	//e.POST("/")

	//register
	//e.POST("/")

	//profile
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, model.Message{
			Message: "Hi",
		})
	})

	routes.InitRoutes(e)

	//Server
	if err := e.Start(":8888"); err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Error to start Server:%v", err)
	}

}

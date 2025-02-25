package main

import (
	"rest/cmd/api/handlers"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/health-check", handlers.HealthCheckHandler)
	e.GET("/posts", handlers.PostIndexHandler)
	e.GET("/post/:id", handlers.PostSingleHandler)
	e.POST("/posts", handlers.HandlePost)
	e.DELETE("/post/:id", handlers.Removeuser)
	e.Logger.Fatal(e.Start(":8080"))

}

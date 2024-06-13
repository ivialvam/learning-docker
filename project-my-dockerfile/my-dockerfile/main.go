package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "github.com/gin-gonic/gin"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK,
			"my-go-app:v2")
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "5555"
	}
	e.Logger.Fatal(e.Start(":" + httpPort))
}

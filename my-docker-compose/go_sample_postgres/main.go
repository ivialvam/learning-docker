package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "github.com/gin-gonic/gin"
)

func main() {
	// Get env
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB HOST")
	dbPort := os.Getenv("DB_PORT")

	// Connect to database
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s? sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println("Connect to: ", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Connection to database failed: ", err)
	}

	// test ping to database
	if err := db.Ping(); err != nil {
		fmt.Printf("Connection to database failed (DB_HOST: %s):%s\n", dbHost, err)
	} else {
		fmt.Println("Successfully connected to the database: ", db)
	}

	// web service
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK,
			"Hellow, Doktahh! <3")
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	e.Logger.Fatal(e.Start(":" + httpPort))
}

package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Get environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Construct the connection string
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println("Connect to: ", connStr)

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}

	// Test the database connection
	if err := db.Ping(); err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
	} else {
		fmt.Println("Successfully connected to database!")
		fmt.Println("Database: ", db)
	}

	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define routes
	e.GET("/", func(c echo.Context) error {
		fmt.Println("To Infinity and Beyond!")
		return c.String(http.StatusOK, "Ouch!")
	})

	e.GET("/ping", func(c echo.Context) error {
		if err := db.Ping(); err != nil {
			fmt.Println("Error pinging database:", err)
			return c.String(http.StatusInternalServerError, "Database connection failed")
		} else {
			fmt.Println("Ping to database successful")
			return c.String(http.StatusOK, "Ping to Database connection successful")
		}
	})

	// Start the server on port 78
	err = e.Start(":78")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

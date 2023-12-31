package main

import (
	"fmt"
	"os"
	"prakerja6/config"
	"prakerja6/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDatabase()
	e := echo.New()
	e = routes.InitRoute(e)
	e.Start(getPort())
}

func getPort() string {
	port := os.Getenv("3306")
	if port == "" {
		port = "8000" // Default port if not specified
	}

	return fmt.Sprintf(":%s", port)
}
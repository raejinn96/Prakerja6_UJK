package main

import "github.com/labstack/echo/v4"

// "fmt"

type Users struct{
	Names []string
	Country []string
	Category []int
	Interest []string
}

func main() {

	e := echo.New()
	
	e.GET("/profile", profile)
	e.Start(":8000")
}

func profile(c echo.Context) error {
	return c.JSON(200, "Good") 
}
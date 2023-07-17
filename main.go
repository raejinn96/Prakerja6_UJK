package main

import "github.com/labstack/echo/v4"

// "fmt"

type Users struct{
	Names string
	Country string
	Category int
	Interest string
}

func main() {

	e := echo.New()
	
	e.GET("/profile", respond)
	e.GET("/users", usersData)

	e.Start(":8000")
}

func usersData(c echo.Context) error {
	var user []Users 
	user = append(user, Users{"Aslew", "Cambodia", 3, "Technology"})
	user = append(user, Users{"Jino", "Singapore", 1, "Entertainment"})

	return c.JSON(200, user) 
}

func respond(c echo.Context) error {
	return c.JSON(200, "Good") 
}
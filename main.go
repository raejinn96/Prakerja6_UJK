package main

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// "fmt"

type Users struct{
	Id int
	Names string
	Country string
	Category int
	Interest string
}

func main() {

	e := echo.New()
	
	// Endpoint untuk GET :
	e.GET("/profile", respond)
	e.GET("/users", usersData)
	e.GET("/users/:id", individual)

	// Endpoint untuk Post :
	e.POST("/users", writeUser)

	e.Start(":8000")
}

func writeUser(c echo.Context) error {
	id,_ := strconv.Atoi(c.FormValue("id"))
	name := c.FormValue("name")
	country := c.FormValue("country")
	category,_ := strconv.Atoi(c.FormValue("category"))
	interest := c.FormValue("interest")

	user := Users{id, name, country, category, interest}

	return c.JSON(200, user) 
}

func individual(c echo.Context) error {
	var id,_ = strconv.Atoi(c.Param("id"))
	var user Users = Users{id, "Aslew", "Cambodia", 3, "Technology"}
	

	return c.JSON(200, user) 
}

func usersData(c echo.Context) error {
	var user []Users 
	user = append(user, Users{1, "Aslew", "Cambodia", 3, "Technology"})
	user = append(user, Users{2, "Jino", "Singapore", 1, "Entertainment"})

	return c.JSON(200, user) 
}

func respond(c echo.Context) error {
	return c.JSON(200, "Good") 
}
package controllers

import (
	"net/http"
	"prakerja6/config"
	"prakerja6/models"
	"strconv"

	"github.com/labstack/echo/v4"
)


func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Message: "Gagal memasukkan data", Status: false, Data: nil,
		})
	}
	return c.JSON(http.StatusCreated, models.Response{
		Message: "Success memasukkan data", Status: true, Data: user,
	})
}

func DetailUserController(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))
	var category, _ = strconv.Atoi(c.Param("category"))
	var user models.User = models.User{Id: id, Name: "Aslew", Country: "Cambodia", Category: category, Interest: "Technology"}
	return c.JSON(200, user)
}

func UserController(c echo.Context) error {
	var users []models.User 
	
	result := config.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Message: "Gagal mendapatkan data", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "Success mendapatkan data", Status: true, Data: users,
	})
}

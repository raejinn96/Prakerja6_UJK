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
}
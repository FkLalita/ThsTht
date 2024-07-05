package main

import (
	"net/http"

	"github.com/FkLalita/ThsTht/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	utils.DecodeJSON()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

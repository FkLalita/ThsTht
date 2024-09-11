package main

import (
	"net/http"

	"github.com/FkLalita/ThsTht/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(e echo.Context) error {
		return e.JSON(http.StatusOK, handlers.Index(e))
	})
	e.Logger.Fatal(e.Start(":8080"))
}

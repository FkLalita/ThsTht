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
	e.GET("/:id/", func(e echo.Context) error {
		return e.JSON(http.StatusOK, handlers.GetQuestion(e))
	})
	e.POST("/question/:id/vote", func(e echo.Context) error {
		return e.JSON(http.StatusOK, handlers.Vote(e))
	})
	_ = handlers.CreateQuestion()
	e.Logger.Fatal(e.Start(":8080"))

}

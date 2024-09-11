package handlers

import (
	//	"github.com/FkLalita/ThsTht/models"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Index(e echo.Context) error {
	data, err := os.Open("models/data.json")
	if err != nil {
		fmt.Println("err opening file:", err)
	}
	return e.JSON(http.StatusOK, data)
}

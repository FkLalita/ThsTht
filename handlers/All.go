package handlers

import (
	//	"github.com/FkLalita/ThsTht/models"
	"net/http"

	"github.com/FkLalita/ThsTht/utils"
	"github.com/labstack/echo/v4"
)

func Index(e echo.Context) error {
	data := utils.DecodeJSON()
	return e.JSON(http.StatusOK, data)
}

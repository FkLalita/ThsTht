package handlers

import (
	"fmt"
	"github.com/FkLalita/ThsTht/models"
	"net/http"

	//	"github.com/FkLalita/ThsTht/utils"
	"github.com/labstack/echo/v4"
)

func Index(e echo.Context) error {
	data, err := models.GetAllQuestions()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	return e.JSON(http.StatusOK, data)
}

func GetQuestion(e echo.Context) error {
	return nil
}

func Vote(e echo.Context) error {
	return nil
}

func CreateQuestion() error {
	var (
		questions_category = "sport"
		questions_text     = "Ronaldo or messi"
		choice1_text       = "Ronaldo"
		choice1_context    = "5 balon dor "
		choice2_text       = "messi"
		choice2_context    = "8 balon dor"
	)
	err := models.CreateQuestion(questions_category, questions_text, choice1_text, choice1_context, choice2_text, choice2_context)
	if err != nil {
		fmt.Println("error in cReation of questions", err)
		return err
	}
	return nil
}

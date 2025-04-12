package models

import (
	//	"github.com/FkLalita/ThsTht/utils"
	"fmt"
	"github.com/FkLalita/ThsTht/db"
	"time"
)

type Question struct {
	ID           int       `json:"id"`
	Category     string    `json:"category"`
	QuestionText string    `json:"question_text"`
	Choice1      Choice    `json:"choice1"`
	Choice2      Choice    `json:"choice2"`
	CreatedAt    time.Time `json:"created_at"`
}

func GetAllQuestions() ([]Question, error) {
	row, err := db.DB.Query("SELECT * FROM questions")
	if err != nil {
		fmt.Println("Error getting data from database:", err)
		return nil, err
	}

	var questions []Question
	for row.Next() {
		var question Question
		err = row.Scan(&question.ID, &question.Category, &question.QuestionText, &question.Choice1.Text, &question.Choice1.Context, &question.Choice1.Count, &question.Choice2.Text, &question.Choice2.Context, &question.Choice2.Count, &question.CreatedAt)
		if err != nil {
			fmt.Println("Error Scanning to variable", err)
			return nil, err
		}
		questions = append(questions, question)
	}
	return questions, nil
}

func CreateQuestion(question_category string, question_text string, choice1_text string, choice1_context string, choice2_text string, choice2_context string) error {
	_, err := db.DB.Exec("INSERT INTO questions (question_category, question_text, choice1_text, choice1_context, choice2_text, choice2_context) VALUES (?,?,?,?,?,?)", question_category, question_text, choice1_text, choice1_context, choice2_text, choice2_context)
	if err != nil {
		fmt.Println("Error creating question:", err)
		return err
	}
	return nil
}

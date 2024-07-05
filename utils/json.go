package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/FkLalita/ThsTht/models"
)

type Questions []models.Question

func DecodeJSON() {
	questions := Questions{}

	file, err := os.Open("models/data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&questions)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Decoded questions:")
	for _, question := range questions {
		fmt.Printf("  Category: %s\n", question.Category)
		fmt.Printf("  Question: %s\n", question.QuestionText)
		fmt.Printf("    Choice 1: %s (%s)\n", question.Choice1, question.Meaning1)
		fmt.Printf("    Choice 2: %s (%s)\n", question.Choice2, question.Meaning2)
	}
}

package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/FkLalita/ThsTht/models"
)

type Questions []models.Question

func DecodeJSON() Questions {
	questions := Questions{}

	file, err := os.Open("models/data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&questions)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	fmt.Println("Decoded questions:")
	return questions
}

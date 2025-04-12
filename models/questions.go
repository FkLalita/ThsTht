package models

import (
// "github.com/FkLalita/ThsTht/utils"
)

type Choice struct {
	Text    string `json:"text"`
	Context string `json:"context"`
	Count   int    `json:"count"`
}

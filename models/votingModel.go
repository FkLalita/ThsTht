package models

import "time"

type Vote struct {
	ID         int       `json:"id"`
	IPAddress  string    `json:"ip_address"`
	QuestionID int       `json:"question_id"`
	Choice     string    `json:"choice"`
	VoteTime   time.Time `json:"vote_time"`
}

type Question struct {
	ID           int    `json:"id"`
	Category     string `json:"category"`
	QuestionText string `json:"question_text"`
	Choice1      string `json:"choice1"`
	Choice2      string `json:"choice2"`
	Meaning1     string `json:"meaning1,omitempty"`
	Meaning2     string `json:"meaning2,omitempty"`
}

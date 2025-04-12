package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func init() {
	var err error

	DB, err = sql.Open("sqlite3", "db/questions.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
	}

	_ = DB.Ping

	// Create questions table
	_, err = DB.Exec(`
		 CREATE TABLE IF NOT EXISTS questions (
			 question_id INTEGER PRIMARY KEY AUTOINCREMENT,
			 question_category VARCHAR(255) NOT NULL,
			 question_text TEXT NOT NULL,

			
			 choice1_text,
			 choice1_context,
			 choice1_votes INTEGAR DEFAULT 0,

			
			 choice2_text TEXT NOT NULL,
			 choice2_context,
			 choice2_votes INTEGER DEFAULT 0,


			 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
 	`)
	if err != nil {
		fmt.Println("error creating table", err)
	}

	fmt.Println("questions table created successfully")

	fmt.Println("Connected............")
}

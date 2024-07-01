package repository

import (
	"database/sql"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error

	dsn := ""

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
}

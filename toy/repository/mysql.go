package repository

import (
	"database/sql"
	"fmt"
	"log"
	"toy/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	cfg := config.AppConfig.MySQL
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
}

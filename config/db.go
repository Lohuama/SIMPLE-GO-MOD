package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func SetupDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	_ = os.Getenv("DB_HOST")
	_ = os.Getenv("DB_PORT")
	_ = os.Getenv("DB_USERNAME")
	_ = os.Getenv("DB_PASSWORD")
	_ = os.Getenv("DB_DATABASE")

	_ = fmt.Sprint("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable")
}
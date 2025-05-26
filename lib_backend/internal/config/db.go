package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	err := godotenv.Load()

	if err != nil {
		log.Println(".env not found")
	}
}

func SetupDB() (*sql.DB, error) {
	LoadEnv()

	host := os.Getenv("HOST_DB")
	port := os.Getenv("PORT_DB")
	user := os.Getenv("USER_DB")
	password := os.Getenv("PASSWORD_DB")
	dbname := os.Getenv("NAME_DB")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	log.Printf("DEBUG: DSN being used: %s", dsn)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

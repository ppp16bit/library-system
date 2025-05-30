package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

/*func LoadEnv() {

	err := godotenv.Load()

	if err != nil {
		log.Println(".env not found")
	}
}*/

func SetupDB() (*sql.DB, error) {
	//LoadEnv()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	log.Printf("DEBUG: Variables read by os.Getenv(): Host='%s', Port='%s', User='%s', Password='%s', DBName='%s'", host, port, user, password, dbname)

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

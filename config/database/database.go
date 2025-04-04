package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Importing the pq driver
	"os"
)

func DB_connect() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: could not load .env file, using environment variables")
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbname := os.Getenv("POSTGRES_DB")

	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	// test the connection

	if err = db.Ping(); err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	fmt.Println("Successfully connected to database")
	return db, nil
}

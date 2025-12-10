package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := getConnectionString()

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic("could not connect to database: " + err.Error())
	}

	if err = DB.Ping(); err != nil {
		panic("could not ping database: " + err.Error())
	}

	fmt.Println("Database connected successfully")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func getConnectionString() string {
	// Use DATABASE_URL if available (production environments)
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		return dbURL
	}

	// Fallback to individual variables (local development)
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
}

func createTable() {
	createJobsTable := `
	CREATE TABLE IF NOT EXISTS jobs (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT NOT NULL
	)`

	if _, err := DB.Exec(createJobsTable); err != nil {
		panic("could not create jobs table: " + err.Error())
	}
}

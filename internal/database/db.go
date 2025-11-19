package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var connStr string

	// Check if DATABASE_URL exists (common for Render, Heroku, etc.)
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		connStr = dbURL
	} else {
		// Build PostgreSQL connection string from individual variables
		host := getEnv("DB_HOST", "localhost")
		port := getEnv("DB_PORT", "5432")
		user := getEnv("DB_USER", "postgres")
		password := getEnv("DB_PASSWORD", "your_password")
		dbname := getEnv("DB_NAME", "jobboard")
		sslmode := getEnv("DB_SSLMODE", "disable")

		connStr = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			host, port, user, password, dbname, sslmode,
		)
	}

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic("could not connect to database: " + err.Error())
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		panic("could not ping database: " + err.Error())
	}

	fmt.Println("✅ Database connected successfully!")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
	fmt.Println("✅ Database tables created successfully!")

}

func createTable() {
	createJobsTable := `
	CREATE TABLE IF NOT EXISTS jobs (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createJobsTable)
	if err != nil {
		panic("could not create jobs table: " + err.Error())
	}
}

func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

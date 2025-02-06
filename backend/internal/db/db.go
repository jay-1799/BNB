package db

import (
	"database/sql"
	"log"
)

func InitDB() *sql.DB {
	// connStr := "postgres://postgres:password@postgres:5432/uptime_db?sslmode=disable"
	connStr := "postgres://myuser:mypassword@postgres:5432/mydatabase?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

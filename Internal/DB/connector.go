package db

import (
	"database/sql"
	"path/filepath"

	_ "modernc.org/sqlite" // Import SQLite3 driver
)

func connectDb(db_path string) (*sql.DB, error) {

	// Function to intiate connection with the database
	// Expects path to database as db_path and returns database connection and error if any

	cleaned_path := filepath.Clean(db_path)

	db, err := sql.Open("sqlite", cleaned_path)
	if err != nil {

		return nil, err

	}

	if err = db.Ping(); err != nil {

		return nil, err

	}
	return db, err

}

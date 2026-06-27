// internal/repository/db_repository.go
package repository

import (
	"database/sql"
	"fmt"
	"MME/internal/domain"
	
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver anonymously
)

type DbRepository struct {
	db *sql.DB
}

// NewDbRepository initializes the database file and creates the table if it does not exist
func NewDbRepository(dbPath string) (*DbRepository, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Create a simple table to store our extracted metadata logs
	query := `
	CREATE TABLE IF NOT EXISTS image_metadata (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		format TEXT,
		width INTEGER,
		height INTEGER,
		size_bytes INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return &DbRepository{db: db}, nil
}

// SaveMetadata inserts a new metadata record into the database
func (r *DbRepository) SaveMetadata(metadata *domain.ImageMetadata) error {
	query := `INSERT INTO image_metadata (format, width, height, size_bytes) VALUES (?, ?, ?, ?)`
	
	_, err := r.db.Exec(query, metadata.Format, metadata.Width, metadata.Height, metadata.SizeBytes)
	if err != nil {
		return fmt.Errorf("failed to insert metadata: %w", err)
	}

	fmt.Println("💾 Successfully saved metadata record to SQLite DB!")
	return nil
}

// Close gracefully terminates the database connection pool
func (r *DbRepository) Close() error {
	return r.db.Close()
}
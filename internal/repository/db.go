package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dataSourceName string) *sql.DB {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	createTables(db)
	return db
}

func createTables(db *sql.DB) {
	objectivesTable := `
	CREATE TABLE IF NOT EXISTS objectives (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		priority INTEGER NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT 0,
		deadline TEXT NOT NULL
	);`

	tasksTable := `
	CREATE TABLE IF NOT EXISTS tasks (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		objective_id TEXT NOT NULL,
		priority INTEGER NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT 0,
		deadline TEXT NOT NULL,
		FOREIGN KEY(objective_id) REFERENCES objectives(id) ON DELETE CASCADE
	);`

	if _, err := db.Exec(objectivesTable); err != nil {
		log.Fatalf("Failed to create objectives table: %v", err)
	}

	if _, err := db.Exec(tasksTable); err != nil {
		log.Fatalf("Failed to create tasks table: %v", err)
	}
}

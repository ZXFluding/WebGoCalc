package postgres

import (
	"database/sql"
	"fmt"
	"test/services/config"

	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func New(storage *config.Storage) (*Repository, error) {
	const op = "repository.postgres.New"
	db, err := sql.Open("postgres", storage.AsString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	// stmt, err := db.Prepare(`
	// CREATE TABLE IF NOT EXISTS students(
	// 	id INTEGER PRIMARY KEY,
	// 	name TEXT NOT NULL UNIQUE,
	// 	firstName CHARACTER VARYING(30),
	// 	lastName CHARACTER VARYING(30),
	// 	scool CHARACTER VARYING(30),
	// 	class CHARACTER VARYING(1),
	// 	day  CHARACTER VARYING(1),
	// 	time INTEGER,
	// 	cost INTEGER);
	// CREATE INDEX IF NOT EXISTS idx_name ON students(alias);
	// `)
	// _ = stmt
	// if err != nil {
	// 	return nil, fmt.Errorf("%s: %w", op, err)
	// }
	return &Repository{db: db}, nil
}
